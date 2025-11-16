package db

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"log"
	"os"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mvt "github.com/paulmach/orb/encoding/mvt"
	wkt "github.com/paulmach/orb/encoding/wkt"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
	"github.com/paulmach/orb/project"
	"github.com/rabbit-backend/go-tiles/models"
	"github.com/rabbit-backend/go-tiles/utils"
	engine "github.com/rabbit-backend/template"
)

func init() {
	godotenv.Load()

	if os.Getenv("MEMSQL_CERT_PATH") != "" {
		cert, err := os.ReadFile(os.Getenv("MEMSQL_CERT_PATH"))
		if err != nil {
			return
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(cert)

		mysql.RegisterTLSConfig("memsql", &tls.Config{RootCAs: caCertPool})
	}

}

type MemSQLSource struct {
	db     *sql.DB
	engine *engine.Engine
}

func (s *MemSQLSource) Open(conn string) {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatalln(err)
	}

	s.db = db
}

func (s *MemSQLSource) Execute(c echo.Context, queryPath string, params any) ([]byte, error) {
	var buf []byte

	req, err := models.MapRequetDecode(params)
	if err != nil {
		return nil, err
	}

	query, args, err := s.engine.Execute(queryPath, params)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data, err := utils.RowToJson(rows)
	if err != nil {
		return nil, err
	}

	// probably the data is list of objects need to iterate over them
	// create geojson features and encode them into mvt tiles
	fc := geojson.NewFeatureCollection()
	for _, item := range data {
		geomWKTByte, ok := item["geom"].([]byte)
		if !ok {
			continue
		}

		geomWKT := string(geomWKTByte)
		geom, err := wkt.Unmarshal(geomWKT)
		if err != nil {
			continue
		}

		feature := geojson.NewFeature(project.Geometry(geom, project.Mercator.ToWGS84))
		feature.Properties = geojson.Properties{}

		for key, value := range item {
			if key == "geom" {
				continue
			}

			feature.Properties[key] = value
		}

		fc.Append(feature)
	}

	layers := mvt.NewLayers(map[string]*geojson.FeatureCollection{
		"data": fc,
	})

	tileForProjection := maptile.New(uint32(req.X), uint32(req.Y), maptile.Zoom(req.Z))
	layers.ProjectToTile(tileForProjection)

	layers.Clip(mvt.MapboxGLDefaultExtentBound)
	buf, err = mvt.Marshal(layers)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

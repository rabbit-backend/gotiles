package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	mvt "github.com/paulmach/orb/encoding/mvt"
	wkt "github.com/paulmach/orb/encoding/wkt"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
	"github.com/rabbit-backend/go-tiles/models"
	"github.com/rabbit-backend/go-tiles/utils"
	engine "github.com/rabbit-backend/template"
)

type MySQLSource struct {
	db     *sql.DB
	engine *engine.Engine
}

func (s *MySQLSource) Open(conn string) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}

	s.db = db
}

func (s *MySQLSource) Execute(c echo.Context, queryPath string, params any) ([]byte, error) {
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
	for id, item := range data {
		geomWKT, ok := item["geom"].(string)
		if !ok {
			continue
		}

		geom, err := wkt.Unmarshal(geomWKT)
		if err != nil {
			continue
		}

		feature := geojson.NewFeature(geom)
		feature.Properties = geojson.Properties{
			"id": id,
		}

		for key, value := range item {
			if key == "geom" {
				continue
			}

			feature.Properties[key] = value
		}

		fc.Append(feature)
	}

	layers := mvt.NewLayers(map[string]*geojson.FeatureCollection{
		req.LayerName: fc,
	})

	tileForProjection := maptile.New(req.X, req.Y, maptile.Zoom(req.Z))
	layers.ProjectToTile(tileForProjection)

	layers.Clip(mvt.MapboxGLDefaultExtentBound)
	buf, err = mvt.MarshalGzipped(layers)
	if err != nil {
		return nil, err
	}

	c.Response().Header().Add("Content-Encoding", "gzip")
	return buf, nil
}

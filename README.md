Sample Config

```json
{
  "sources": [
    {
      "name": "db_a",
      "connection": {
        "type": "env",
        "value": "POSTGRES_URL"
      }
    },
    {
      "name": "db_b",
      "connection": {
        "type": "env",
        "value": "POSTGRES_URL"
      }
    }
  ]
}
```

Sample Folder Structure

```bash
tiles/
├─ db/
│  ├─ db_a/
│  │  ├─ buildings.sql
│  ├─ db_b/
│  │  ├─ airports.sql
├─ static/
│  ├─ roads.pmtiles

```

Sample SQL query

```sql
SELECT ST_AsMVT(tile, 'buildings', 4096, 'geom') FROM (
    SELECT
        ST_AsMVTGeom(
            ST_Transform(ST_CurveToLine(geom), 3857),
            ST_TileEnvelope($3, $1, $2),
            4096, 64, true
        ) as geom
    FROM "public"."buildings"
    WHERE geom && ST_Transform(ST_TileEnvelope($3, $1, $2), 4326)
) as tile WHERE geom IS NOT NULL;
```

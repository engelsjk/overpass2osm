# overpass2josm

The overpass2osm package converts Overpass JSON data into an OSM object. Specifically, it's intended to supplement features of the [paulmach/osm](https://github.com/paulmach/osm) package used for handing OpenStreetMap data in Go.

## Example

```go
b := []byte(`
  {
    "version": 0.6,
    "generator": "Overpass API 0.7.56.7 b85c4387",
    "elements": [
      {
        "type": "node",
        "id": 7495431894,
        "lat": 45.7042173,
        "lon": -121.5236360,
        "tags": {
          "emergency": "fire_hydrant",
          "water_source": "main"
        }
      }
    ]
  }
`)

ovp := &overpass2osm.Overpass{}

err := json.Unmarshal(b, &ovp)
if err != nil {
  panic(err)
}

o, err := overpass2osm.Convert(ovp) // o object is type *osm.OSM
if err != nil {
  panic(err)
}
```

The ```*osm.OSM``` object can now be converted to GeoJSON using the sub-package  [paulmach/osm/osmgeojson](https://github.com/paulmach/osm/tree/master/osmgeojson).

```go
fc, err := osmgeojson.Convert(o)
if err != nil {
  panic(err)
}

gj, _ := json.MarshalIndent(fc, "", " ")

fmt.Println(string(gj))
```

The resulting GeoJSON string should now be a direct representation of the Overpass JSON element data. 

```bash
{
 "type": "FeatureCollection",
 "features": [
  {
   "id": "node/7495431894",
   "type": "Feature",
   "geometry": {
    "type": "Point",
    "coordinates": [
     -121.523636,
     45.7042173
    ]
   },
   "properties": {
    "id": 7495431894,
    "meta": {},
    "relations": [],
    "tags": {
     "emergency": "fire_hydrant",
     "water_source": "main"
    },
    "type": "node"
   }
  }
 ]
}

```

## Detail

This overpass2osm package utilizes the [paulmach/osm/annotate](https://github.com/paulmach/osm/tree/master/annotate) sub-package to convert the Overpass type defined in this package into an osm.OSM object. It includes the same options available in [paulmach/osm/annotate](https://github.com/paulmach/osm/tree/master/annotate).

package overpass2osm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/paulmach/osm/osmgeojson"
)

func TestConvert(t *testing.T) {
	ovp := loadOverpass(t, "testdata/overpass.json")
	o, err := Convert(ovp, IgnoreMissingChildren(true))
	if err != nil {
		t.Fatalf("unable to convert overpass to osm: %v", err)
	}

	fc, err := osmgeojson.Convert(o, osmgeojson.NoMeta(true), osmgeojson.NoRelationMembership(true))
	if err != nil {
		t.Fatalf("unable to marshal osm to json: %v", err)
	}

	gj, _ := json.MarshalIndent(fc, "", " ")
	fmt.Println(string(gj))

	// TODO: validate resulting GeoJSON against expected string

	// _ = ioutil.WriteFile("out.geojson", gj, 0644)
}

func loadOverpass(t testing.TB, filename string) *Overpass {
	data := readFile(t, filename)

	ovp := &Overpass{}
	err := json.Unmarshal(data, &ovp)
	if err != nil {
		t.Fatalf("unable to unmarshal %s: %v", filename, err)
	}

	return ovp
}

func readFile(t testing.TB, filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("unable to open %s: %v", filename, err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("unable to read file %s: %v", filename, err)
	}

	return data
}

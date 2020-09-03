package ovptogeojson

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	ovp := loadOverpass(t, "testdata/overpass.json")
	_ = ovp
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

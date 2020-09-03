package overpass2osm

import "encoding/json"

type Overpass struct {
	Version   float64 `json:"version,attr,omitempty"`
	Generator string  `json:"generator,attr,omitempty"`

	OSM3S struct {
		TimeStampOSMBase string `json:"timestamp_osm_base"`
		Copyright        string `json:"copyright"`
	}

	Elements []Elements `json:"elements"`
}

type Elements struct {
	Type    string                   `json:"type"`
	ID      int                      `json:"id"`
	NodeIDS []int                    `json:"nodes"`
	WayIDs  []int                    `json:"ways"`
	Tags    map[string]string        `json:"tags"`
	Lat     float64                  `json:"lat"`
	Lon     float64                  `json:"lon"`
	Members []map[string]interface{} `json:"members`
}

func UnmarshalOverpass(data []byte) (*Overpass, error) {

	ovp := &Overpass{}
	err := json.Unmarshal(data, ovp)
	if err != nil {
		return nil, err
	}

	return ovp, nil
}

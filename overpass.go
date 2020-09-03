package overpass2osm

import "encoding/json"

// Overpass represents the response from an Overpass query (JSON only).
type Overpass struct {
	Version   float64 `json:"version,attr,omitempty"`
	Generator string  `json:"generator,attr,omitempty"`

	OSM3S struct {
		TimeStampOSMBase string `json:"timestamp_osm_base,attr,omitempty"`
		Copyright        string `json:"copyright,attr,omitempty"`
	}

	Elements []Element `json:"elements,attr,omitempty"`
}

// Element represents a generalized OSM type (node, way or relation)
// that's listed in an Overpass response.
type Element struct {
	Type    string  `json:"type,attr,omitempty"`
	ID      int     `json:"id,attr,omitempty"`
	Nodes   []int   `json:"nodes,attr,omitempty"`
	Ways    []int   `json:"ways,attr,omitempty"`
	Tags    Tags    `json:"tags,attr,omitempty"`
	Lat     float64 `json:"lat,attr,omitempty"`
	Lon     float64 `json:"lon,attr,omitempty"`
	Members Members `json:"members,attr,omitempty`
}

// Tags represents a map of key/value tag strings.
type Tags map[string]string

// Members represents a set of relation members from an Overpass response element.
type Members []Member

// Member represents a single relation member.
type Member struct {
	Type string `json:"type,attr,omitempty"`
	Ref  int    `json:"ref,attr,omitempty"`
	Role string `json:"role,attr,omitempty"`
}

// UnmarshalOverpass will unmarshal the JSON data into an Overpass object.
func UnmarshalOverpass(data []byte) (*Overpass, error) {

	ovp := &Overpass{}
	err := json.Unmarshal(data, ovp)
	if err != nil {
		return nil, err
	}

	return ovp, nil
}

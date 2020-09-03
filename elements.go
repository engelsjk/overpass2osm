package overpass2osm

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

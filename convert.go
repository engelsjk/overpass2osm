package overpass2osm

import (
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/osm"
)

type context struct {
	ovp       *Overpass
	osm       *osm.OSM
}

// Convert takes a set of osm elements and converts them
// to a geojson feature collection.
func Convert(o *Overpass) (*osm.OSM, error) {

	ctx := &context{
		ovp: o,
		osm: &osm.OSM{},
	}

	for _, elem := range ctx.ovp.Elements {

		// nodes
		if elem.Type == string(osm.TypeNode) {
			node := &osm.Node{
				ID:      osm.NodeID(elem.ID),
				Lat:     elem.Lat,
				Lon:     elem.Lon,
				Tags:    unmapTags(elem.Tags),
				Visible: true,
			}
			ctx.osm.Nodes = append(ctx.osm.Nodes, node)
		}

		// ways
		if elem.Type == string(osm.TypeWay) {
			way := &osm.Way{
				ID:      osm.WayID(elem.ID),
				Tags:    unmapTags(elem.Tags),
				Visible: true,
			}
			ctx.osm.Ways = append(ctx.osm.Ways, way)
		}

		// relations
		if elem.Type == string(osm.TypeRelation) {
			relation := &osm.Relation{
				ID:      osm.WayID(elem.ID),
				Tags:    unmapTags(elem.Tags),
				Members: unmapMembers(elem.Members),
				Visible: true,
			}
			ctx.osm.Ways = append(ctx.osm.Ways, way)
		}
	}

	ds := ctx.osm.HistoryDatasource()

	err = annotate.Ways(ctx, ctx.osm.Ways, ds)
	if err != nil {
		return nil, err
	}

	err = annotate.Relations(ctx, ctx.osm.Relations, ds)
	if err != nil {
		return nil, err
	}

	return ctx.osm, nil
}

func unmapTags(t map[string]string) osm.Tags {
	tags := osm.Tags{}
	for k, v := range t {
		tag := osm.Tags{Key: k, Value: v}
		tags = append(tags, tag)
	}
	return tags
}

func unmapMembers(members []map[string]interface{}) osm.Members {
	members := osm.Members{}
	for _, m := range members {
		member := osm.Members{
			Type: osm.Type(m["type"]), 
			Ref: int64(m["ref"]),
			Role: m[["role"]]
		}
		members = append(members,member)
	}
	return members
}

package overpass2osm

import (
	"context"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/annotate"
)

// Convert takes an Overpass object and converts it to an osm.OSM object.
func Convert(ovp *Overpass, opts ...annotate.Option) (*osm.OSM, error) {

	o := &osm.OSM{}

	for _, elem := range ovp.Elements {

		// nodes
		if elem.Type == string(osm.TypeNode) {
			node := &osm.Node{
				ID:      osm.NodeID(elem.ID),
				Lat:     elem.Lat,
				Lon:     elem.Lon,
				Tags:    unmapTags(elem.Tags),
				Visible: true,
			}
			o.Nodes = append(o.Nodes, node)
		}

		// ways
		if elem.Type == string(osm.TypeWay) {
			way := &osm.Way{
				ID:      osm.WayID(elem.ID),
				Tags:    unmapTags(elem.Tags),
				Nodes:   unmapWayNodes(elem.Nodes),
				Visible: true,
			}
			o.Ways = append(o.Ways, way)
		}

		// relations
		if elem.Type == string(osm.TypeRelation) {
			relation := &osm.Relation{
				ID:      osm.RelationID(elem.ID),
				Tags:    unmapTags(elem.Tags),
				Members: unmapRelationMembers(elem.Members),
				Visible: true,
			}
			o.Relations = append(o.Relations, relation)
		}
	}

	ctx := context.Background()

	ds := o.HistoryDatasource()

	err := annotate.Ways(ctx, o.Ways, ds, opts...)
	if err != nil {
		return nil, err
	}

	err = annotate.Relations(ctx, o.Relations, ds, opts...)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func unmapTags(ts map[string]string) osm.Tags {
	tags := osm.Tags{}
	for k, v := range ts {
		tag := osm.Tag{Key: k, Value: v}
		tags = append(tags, tag)
	}
	return tags
}

func unmapWayNodes(nodes []int) osm.WayNodes {
	wayNodes := osm.WayNodes{}
	for _, n := range nodes {
		wayNodes = append(wayNodes, osm.WayNode{ID: osm.NodeID(n)})
	}
	return wayNodes
}

func unmapRelationMembers(ms Members) osm.Members {
	members := osm.Members{}
	for _, m := range ms {
		member := osm.Member{
			Type: osm.Type(m.Type),
			Ref:  int64(m.Ref),
			Role: m.Role,
		}
		members = append(members, member)
	}
	return members
}

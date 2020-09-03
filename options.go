package overpass2osm

import (
	"context"
	"time"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/annotate"
)

type Option func(ctx context.Context) error

const defaultThreshold = 30 * time.Minute

func Threshold(t time.Duration) annotate.Option {
	return annotate.Threshold(t)
}

func IgnoreInconsistency(yes bool) annotate.Option {
	return annotate.IgnoreInconsistency(yes)
}

func IgnoreMissingChildren(yes bool) annotate.Option {
	return annotate.IgnoreMissingChildren(yes)
}

func ChildFilter(filter func(osm.FeatureID) bool) annotate.Option {
	return annotate.ChildFilter(filter)
}

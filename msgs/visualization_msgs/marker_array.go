package visualization_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type MarkerArray struct {
	msgs.Package `ros:"visualization_msgs"`
	Markers      []Marker
}

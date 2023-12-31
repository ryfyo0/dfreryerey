//autogenerated:yes
//nolint:revive,lll
package geographic_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/uuid_msgs"
)

type GeographicMap struct {
	msg.Package `ros:"geographic_msgs"`
	Header      std_msgs.Header
	Id          uuid_msgs.UniqueID
	Bounds      BoundingBox
	Points      []WayPoint
	Features    []MapFeature
	Props       []KeyValue
}

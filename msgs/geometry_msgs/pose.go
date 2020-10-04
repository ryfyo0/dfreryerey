package geometry_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Pose struct {
	msgs.Package `ros:"geometry_msgs"`
	Position     Point
	Orientation  Quaternion
}

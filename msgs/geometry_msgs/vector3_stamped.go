package geometry_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type Vector3Stamped struct {
	msgs.Package `ros:"geometry_msgs"`
	Header       std_msgs.Header
	Vector       Vector3
}

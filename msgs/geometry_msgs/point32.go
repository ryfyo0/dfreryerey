package geometry_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Point32 struct {
	msgs.Package `ros:"geometry_msgs"`
	X            float32
	Y            float32
	Z            float32
}

package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type MultiArrayDimension struct {
	msgs.Package `ros:"std_msgs"`
	Label        string
	Size         uint32
	Stride       uint32
}

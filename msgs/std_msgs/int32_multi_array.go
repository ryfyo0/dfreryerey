package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Int32MultiArray struct {
	msgs.Package `ros:"std_msgs"`
	Layout       MultiArrayLayout
	Data         []int32
}

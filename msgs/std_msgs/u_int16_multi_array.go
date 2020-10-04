package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type UInt16MultiArray struct {
	msgs.Package `ros:"std_msgs"`
	Layout       MultiArrayLayout
	Data         []uint16
}

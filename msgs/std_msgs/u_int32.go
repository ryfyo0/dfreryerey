package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type UInt32 struct {
	msgs.Package `ros:"std_msgs"`
	Data         uint32
}

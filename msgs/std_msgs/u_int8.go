package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type UInt8 struct {
	msgs.Package `ros:"std_msgs"`
	Data         uint8
}

package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Int32 struct {
	msgs.Package `ros:"std_msgs"`
	Data         int32
}

package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Byte struct {
	msgs.Package `ros:"std_msgs"`
	Data         int8 `ros:"byte"`
}

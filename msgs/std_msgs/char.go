package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Char struct {
	msgs.Package `ros:"std_msgs"`
	Data         uint8 `ros:"char"`
}

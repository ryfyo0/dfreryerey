package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Bool struct {
	msgs.Package `ros:"std_msgs"`
	Data         bool
}

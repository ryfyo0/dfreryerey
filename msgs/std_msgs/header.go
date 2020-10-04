package std_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"time"
)

type Header struct {
	msgs.Package `ros:"std_msgs"`
	Seq          uint32
	Stamp        time.Time
	FrameId      string
}

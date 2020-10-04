package visualization_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
	"time"
)

type ImageMarker struct {
	msgs.Package     `ros:"visualization_msgs"`
	msgs.Definitions `ros:"uint8 CIRCLE=0,uint8 LINE_STRIP=1,uint8 LINE_LIST=2,uint8 POLYGON=3,uint8 POINTS=4,uint8 ADD=0,uint8 REMOVE=1"`
	Header           std_msgs.Header
	Ns               string
	Id               int32
	Type             int32
	Action           int32
	Position         geometry_msgs.Point
	Scale            float32
	OutlineColor     std_msgs.ColorRGBA
	Filled           uint8
	FillColor        std_msgs.ColorRGBA
	Lifetime         time.Duration
	Points           []geometry_msgs.Point
	OutlineColors    []std_msgs.ColorRGBA
}

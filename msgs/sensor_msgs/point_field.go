package sensor_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type PointField struct {
	msgs.Package     `ros:"sensor_msgs"`
	msgs.Definitions `ros:"uint8 INT8=1,uint8 UINT8=2,uint8 INT16=3,uint8 UINT16=4,uint8 INT32=5,uint8 UINT32=6,uint8 FLOAT32=7,uint8 FLOAT64=8"`
	Name             string
	Offset           uint32
	Datatype         uint8
	Count            uint32
}

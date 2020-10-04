package sensor_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type LaserEcho struct {
	msgs.Package `ros:"sensor_msgs"`
	Echoes       []float32
}

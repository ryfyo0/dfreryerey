package geometry_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type AccelWithCovariance struct {
	msgs.Package `ros:"geometry_msgs"`
	Accel        Accel
	Covariance   [36]float64
}

package shape_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Plane struct {
	msgs.Package `ros:"shape_msgs"`
	Coef         [4]float64
}

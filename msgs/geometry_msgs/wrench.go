package geometry_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type Wrench struct {
	msgs.Package `ros:"geometry_msgs"`
	Force        Vector3
	Torque       Vector3
}

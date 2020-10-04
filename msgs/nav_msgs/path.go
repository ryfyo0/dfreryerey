package nav_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type Path struct {
	msgs.Package `ros:"nav_msgs"`
	Header       std_msgs.Header
	Poses        []geometry_msgs.PoseStamped
}

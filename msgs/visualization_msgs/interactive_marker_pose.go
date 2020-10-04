package visualization_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type InteractiveMarkerPose struct {
	msgs.Package `ros:"visualization_msgs"`
	Header       std_msgs.Header
	Pose         geometry_msgs.Pose
	Name         string
}

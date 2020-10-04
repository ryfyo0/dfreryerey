package shape_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
)

type Mesh struct {
	msgs.Package `ros:"shape_msgs"`
	Triangles    []MeshTriangle
	Vertices     []geometry_msgs.Point
}

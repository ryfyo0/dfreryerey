package shape_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type MeshTriangle struct {
	msgs.Package  `ros:"shape_msgs"`
	VertexIndices [3]uint32
}

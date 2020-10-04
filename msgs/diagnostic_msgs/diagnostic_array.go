package diagnostic_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type DiagnosticArray struct {
	msgs.Package `ros:"diagnostic_msgs"`
	Header       std_msgs.Header
	Status       []DiagnosticStatus
}

//autogenerated:yes
//nolint:revive,lll
package sensor_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type JoyFeedbackArray struct {
	msg.Package `ros:"sensor_msgs"`
	Array       []JoyFeedback
}

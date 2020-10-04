package sensor_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type JoyFeedback struct {
	msgs.Package     `ros:"sensor_msgs"`
	msgs.Definitions `ros:"uint8 TYPE_LED=0,uint8 TYPE_RUMBLE=1,uint8 TYPE_BUZZER=2"`
	Type             uint8
	Id               uint8
	Intensity        float32
}

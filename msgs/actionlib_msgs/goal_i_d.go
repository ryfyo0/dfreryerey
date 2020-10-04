package actionlib_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"time"
)

type GoalID struct {
	msgs.Package `ros:"actionlib_msgs"`
	Stamp        time.Time
	Id           string
}

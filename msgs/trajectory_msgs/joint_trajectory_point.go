package trajectory_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"time"
)

type JointTrajectoryPoint struct {
	msgs.Package  `ros:"trajectory_msgs"`
	Positions     []float64
	Velocities    []float64
	Accelerations []float64
	Effort        []float64
	TimeFromStart time.Duration
}

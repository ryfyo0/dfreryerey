package sensor_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type Imu struct {
	msgs.Package                 `ros:"sensor_msgs"`
	Header                       std_msgs.Header
	Orientation                  geometry_msgs.Quaternion
	OrientationCovariance        [9]float64
	AngularVelocity              geometry_msgs.Vector3
	AngularVelocityCovariance    [9]float64
	LinearAcceleration           geometry_msgs.Vector3
	LinearAccelerationCovariance [9]float64
}

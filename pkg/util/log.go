package util

import (
	"time"

	"github.com/sirupsen/logrus"
)

var (
	layout = "02-01-2006 15:04:05"
)

// LogStart to display the version and the time start of service
func LogStart(version string, timeStart time.Time) {
	logrus.Infof("Service started at %s with version %s", version, timeStart.Format(layout))
}

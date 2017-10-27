package model

import (
	"time"
)

type BatteryStats struct {
	Usage   time.Duration
	Standby time.Duration
}

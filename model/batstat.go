package model

import (
	"time"
)

type BatteryStats struct {
	Usage   time.Duration `json:"Usage"`
	Standby time.Duration `json:Standby`
}

package util

import (
	"time"
)

// NowUnixMs returns the current unix time in milliseconds
func NowUnixMs() int64 {
	return time.Now().Round(time.Millisecond).UnixNano() / 1e6
}

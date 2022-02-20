package counter

import "time"

type Time interface {
	Now() time.Time
}

type RealTime struct{}

func (t RealTime) Now() time.Time {
	return time.Now()
}

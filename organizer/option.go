package organizer

import (
	"time"
)

// Option is option for action.Do
type Option struct {
	mTimeLimit time.Time
}

// GetMTimeLimit get mTimeLimit of option
func (option *Option) GetMTimeLimit() time.Time {
	return option.mTimeLimit
}

// SetMTimeLimit set mTimeLimit of option
func (option *Option) SetMTimeLimit(value time.Time) *Option {
	option.mTimeLimit = value
	return option
}

// SetMTimeLimitToNow set mTimeLimit of option into current time
func (option *Option) SetMTimeLimitToNow() *Option {
	option.mTimeLimit = time.Now()
	return option
}

// NewOption create option
func NewOption() *Option {
	return &Option{
		mTimeLimit: time.Time{},
	}
}

package action

import (
	"time"
)

// Option is option for action.Do
type Option struct {
	mTimeLimit    time.Time
	performPre    bool
	performPost   bool
	performAction bool
}

// GetMTimeLimit get MTime of option
func (option *Option) GetMTimeLimit() time.Time {
	return option.mTimeLimit
}

// SetMTimeLimit set MTime of option
func (option *Option) SetMTimeLimit(value time.Time) *Option {
	option.mTimeLimit = value
	return option
}

// GetPerformAction get performAction of option
func (option *Option) GetPerformAction() bool {
	return option.performAction
}

// SetPerformAction set performAction of option
func (option *Option) SetPerformAction(value bool) *Option {
	option.performAction = value
	return option
}

// GetPerformPre get performPre of option
func (option *Option) GetPerformPre() bool {
	return option.performPre
}

// SetPerformPre set performPre of option
func (option *Option) SetPerformPre(value bool) *Option {
	option.performPre = value
	return option
}

// GetPerformPost get performPost of option
func (option *Option) GetPerformPost() bool {
	return option.performPost
}

// SetPerformPost set performPost of option
func (option *Option) SetPerformPost(value bool) *Option {
	option.performPost = value
	return option
}

// NewOption create option
func NewOption() *Option {
	return &Option{
		mTimeLimit:    time.Time{},
		performPre:    true,
		performPost:   true,
		performAction: true,
	}
}

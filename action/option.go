package action

import (
	"path/filepath"
	"time"
)

// Option is option for action.Do
type Option struct {
	mTimeLimit           time.Time
	isPerformPre         bool
	isPerformPost        bool
	isPerformAction      bool
	isRecursiveWorkDir   bool
	workDir              string
	isRecursiveScriptDir bool
	scriptDir            string
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

// GetIsPerformAction get isPerformAction of option
func (option *Option) GetIsPerformAction() bool {
	return option.isPerformAction
}

// SetIsPerformAction setisPperformAction of option
func (option *Option) SetIsPerformAction(value bool) *Option {
	option.isPerformAction = value
	return option
}

// GetIsPerformPre getisPperformPre of option
func (option *Option) GetIsPerformPre() bool {
	return option.isPerformPre
}

// SetIsPerformPre set isPperformPre of option
func (option *Option) SetIsPerformPre(value bool) *Option {
	option.isPerformPre = value
	return option
}

// GetIsPerformPost get isPerformPost of option
func (option *Option) GetIsPerformPost() bool {
	return option.isPerformPost
}

// SetIsPerformPost set isPerformPost of option
func (option *Option) SetIsPerformPost(value bool) *Option {
	option.isPerformPost = value
	return option
}

// GetIsRecursiveWorkDir get isRecursiveWorkDir of option
func (option *Option) GetIsRecursiveWorkDir() bool {
	return option.isRecursiveWorkDir
}

// SetIsRecursiveWorkDir set isRecursiveWorkDir of option
func (option *Option) SetIsRecursiveWorkDir(value bool) *Option {
	option.isRecursiveWorkDir = value
	return option
}

// GetIsRecursiveScriptDir get isRecursiveScriptDir of option
func (option *Option) GetIsRecursiveScriptDir() bool {
	return option.isRecursiveScriptDir
}

// SetIsRecursiveScriptDir set isRecursiveScriptDir of option
func (option *Option) SetIsRecursiveScriptDir(value bool) *Option {
	option.isRecursiveScriptDir = value
	return option
}

// GetWorkDir get workDir of option
func (option *Option) GetWorkDir() string {
	return option.workDir
}

// SetWorkDir set workDir of option
func (option *Option) SetWorkDir(value string) (*Option, error) {
	var err error
	option.workDir, err = filepath.Abs(value)
	return option, err
}

// GetScriptDir get scriptDir of option
func (option *Option) GetScriptDir() string {
	return option.scriptDir
}

// SetScriptDir set scriptDir of option. Also set workDir if it is not set yet
func (option *Option) SetScriptDir(value string) (*Option, error) {
	var err error
	oldScriptDir := option.GetScriptDir()
	option.scriptDir, err = filepath.Abs(value)
	if err != nil {
		return option, err
	}
	if option.GetWorkDir() == oldScriptDir {
		option, err = option.SetWorkDir(value)
	}
	return option, err
}

// NewOption create option
func NewOption() *Option {
	return &Option{
		mTimeLimit:         time.Time{},
		isPerformPre:       true,
		isPerformPost:      true,
		isPerformAction:    true,
		isRecursiveWorkDir: true,
		workDir:            "",
		scriptDir:          "",
	}
}

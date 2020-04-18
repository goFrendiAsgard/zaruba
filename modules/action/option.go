package action

// Option is option for action.Do
// if option.workDir == option.scriptDir, we will run the matching scripts found in every workDir's sub-directory,
// otherwise we will run a matching script in scriptDir into every workDir's sub-directory
type Option struct {
	isPerformPre       bool
	isPerformPost      bool
	isPerformAction    bool
	isRecursiveWorkDir bool
	workDir            string
	scriptDir          string
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

// GetWorkDir get workDir of option
func (option *Option) GetWorkDir() string {
	return option.workDir
}

// SetWorkDir set workDir of option
func (option *Option) SetWorkDir(value string) *Option {
	option.workDir = value
	return option
}

// GetScriptDir get scriptDir of option
func (option *Option) GetScriptDir() string {
	return option.scriptDir
}

// SetScriptDir set scriptDir of option. Also set workDir if it is not set yet
func (option *Option) SetScriptDir(value string) *Option {
	option.scriptDir = value
	return option
}

// NewOption create option
func NewOption() *Option {
	return &Option{
		isPerformPre:       true,
		isPerformPost:      true,
		isPerformAction:    true,
		isRecursiveWorkDir: true,
		workDir:            "",
		scriptDir:          "",
	}
}

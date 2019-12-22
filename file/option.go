package file

// Option is configuration used for second parameter of GetAllFiles
type Option struct {
	maxDepth int
	onlyDir  bool
}

// GetMaxDepth get MTime of option
func (option *Option) GetMaxDepth() int {
	return option.maxDepth
}

// SetMaxDepth set MTime of option
func (option *Option) SetMaxDepth(value int) *Option {
	option.maxDepth = value
	return option
}

// GetOnlyDir get performAction of option
func (option *Option) GetOnlyDir() bool {
	return option.onlyDir
}

// SetOnlyDir set performAction of option
func (option *Option) SetOnlyDir(value bool) *Option {
	option.onlyDir = value
	return option
}

// NewOption get default Option
func NewOption() *Option {
	return &Option{
		maxDepth: 100,
		onlyDir:  false,
	}
}

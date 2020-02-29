package file

// Option is configuration used for second parameter of GetAllFiles
type Option struct {
	maxDepth  int
	isOnlyDir bool
	ignores   []string
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

// GetIsOnlyDir get performAction of option
func (option *Option) GetIsOnlyDir() bool {
	return option.isOnlyDir
}

// SetIsOnlyDir set performAction of option
func (option *Option) SetIsOnlyDir(value bool) *Option {
	option.isOnlyDir = value
	return option
}

// GetIgnores get ignores of option
func (option *Option) GetIgnores() []string {
	return option.ignores
}

// SetIgnores set ignores of option
func (option *Option) SetIgnores(value []string) *Option {
	option.ignores = value
	return option
}

// NewOption get default Option
func NewOption() *Option {
	return &Option{
		maxDepth:  100,
		isOnlyDir: false,
	}
}

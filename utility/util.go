package utility

type Util struct {
	Str *StrUtil
}

func NewUtil() *Util {
	util := &Util{}
	util.Str = NewStrUtil(util)
	return util
}

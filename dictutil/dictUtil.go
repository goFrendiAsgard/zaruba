package dictutil

type DictUtil struct{}

func NewDictUtil() *DictUtil {
	return &DictUtil{}
}

func (dictUtil *DictUtil) GetSortedKeys(dict interface{}) (sortedKeys []string, err error) {
	return DictGetSortedKeys(dict)
}

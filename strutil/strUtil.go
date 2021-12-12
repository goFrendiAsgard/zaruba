package strutil

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type StrUtil struct {
	nameGenerator *NameGenerator
}

func NewStrUtil() *StrUtil {
	return &StrUtil{
		nameGenerator: NewNameGenerator(),
	}
}

func (strUtil *StrUtil) IsUpper(s string) (result bool) {
	return strings.ToUpper(s) == s
}

func (strUtil *StrUtil) IsLower(s string) (result bool) {
	return strings.ToLower(s) == s
}

func (strUtil *StrUtil) ToUpper(s string) (result string) {
	return strings.ToUpper(s)
}

func (strUtil *StrUtil) ToLower(s string) (result string) {
	return strings.ToLower(s)
}

func (strUtil *StrUtil) ToCamel(s string) (result string) {
	rex := regexp.MustCompile("[^a-zA-Z0-9]+")
	strippedStr := rex.ReplaceAllString(s, " ")
	titledStr := strings.Title(strippedStr)
	result = strings.ReplaceAll(titledStr, " ", "")
	if len(result) > 0 {
		firstLetter := strings.ToLower(string(result[0]))
		rest := result[1:]
		return firstLetter + rest
	}
	return result
}

func (strUtil *StrUtil) ToPascal(s string) (result string) {
	return strings.Title(strUtil.ToCamel(s))
}

func (strUtil *StrUtil) ToUpperSnake(s string) (result string) {
	return strUtil.ToUpper(strUtil.ToSnake(s))
}

func (strUtil *StrUtil) ToSnake(s string) (result string) {
	result = ""
	for index, ch := range strUtil.ToCamel(s) {
		if index == 0 {
			result += strings.ToLower(string(ch))
			continue
		}
		if strUtil.IsUpper(string(ch)) {
			result += "_"
		}
		result += strings.ToLower(string(ch))
	}
	return result
}

func (strUtil *StrUtil) ToKebab(s string) (result string) {
	result = ""
	for index, ch := range strUtil.ToCamel(s) {
		if index == 0 {
			result += strings.ToLower(string(ch))
			continue
		}
		if strUtil.IsUpper(string(ch)) {
			result += "-"
		}
		result += strings.ToLower(string(ch))
	}
	return result
}

func (strUtil *StrUtil) Quote(s string, quote byte) (result string) {
	return StrQuote(s, quote)
}

func (strUtil *StrUtil) DoubleQuote(s string) (result string) {
	return StrDoubleQuote(s)
}

func (strUtil *StrUtil) SingleQuote(s string) (result string) {
	return StrSingleQuote(s)
}

func (strUtil *StrUtil) EscapeShellValue(s string) (result string) {
	return StrEscapeShellValue(s)
}

// indent second-last lines
func (strUtil *StrUtil) Indent(multiLineStr string, indentation string) (indentedStr string) {
	return StrIndent(multiLineStr, indentation)
}

func (strUtil *StrUtil) FullIndent(multiLineStr string, indentation string) (indentedStr string) {
	return StrFullIndent(multiLineStr, indentation)
}

func (strUtil *StrUtil) Repeat(s string, repetition int) (result string) {
	return StrRepeat(s, repetition)
}

func (strUtil *StrUtil) GetIndentation(s string, level int) (result string, err error) {
	rex := regexp.MustCompile("^([ \t]+).*$")
	match := rex.FindStringSubmatch(s)
	if len(match) < 2 {
		return "", nil
	}
	totalIndentation := match[1]
	indentationLength := len(totalIndentation) / level
	result = s[:indentationLength]
	if strUtil.Repeat(result, level) != totalIndentation {
		return result, fmt.Errorf("cannot determine single %d indentation for '%s'", level, s)
	}
	return result, nil
}

func (strUtil *StrUtil) Submatch(s string, pattern string) (result []string, err error) {
	return StrSubmatch(s, pattern)
}

func (strUtil *StrUtil) Split(s string, separator string) (result []string) {
	return strings.Split(s, separator)
}

func (strUtil *StrUtil) PadRight(s string, length int, pad string) (result string) {
	for len(s) < length {
		s = s + pad
	}
	return s
}

func (strUtil *StrUtil) PadLeft(s string, length int, pad string) (result string) {
	for len(s) < length {
		s = pad + s
	}
	return s
}

func (struUtil *StrUtil) NewUUID() (uuidStr string) {
	return uuid.NewString()
}

func (struUtil *StrUtil) AddPrefix(s, prefix string) (prefixedStr string) {
	if strings.HasPrefix(s, prefix) {
		return s
	}
	return prefix + s
}

func (strUtil *StrUtil) Trim(str, cutset string) (trimmedStr string) {
	return strings.Trim(str, cutset)
}

func (strUtil *StrUtil) NewName() (randomName string) {
	return strUtil.nameGenerator.Generate()
}

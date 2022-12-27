package strutil

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/google/uuid"
)

type StrUtil struct {
	nameGenerator *NameGenerator
	pluralize     *pluralize.Client
}

func NewStrUtil() *StrUtil {
	return &StrUtil{
		nameGenerator: NewNameGenerator(),
		pluralize:     pluralize.NewClient(),
	}
}

func (strUtil *StrUtil) IsSingular(s string) (result bool) {
	return strUtil.pluralize.IsSingular(s)
}

func (strUtil *StrUtil) IsPlural(s string) (result bool) {
	return strUtil.pluralize.IsPlural(s)
}

func (strUtil *StrUtil) IsReallyUpper(s string) (result bool) {
	// this will only yield true if string is actual upper case letter, not numeric
	return strUtil.IsUpper(s) && !strUtil.IsLower(s)
}

func (strUtil *StrUtil) IsReallyLower(s string) (result bool) {
	// this will only yield true if string is actual lower case letter, not numeric
	return !strUtil.IsUpper(s) && strUtil.IsLower(s)
}

func (strUtil *StrUtil) IsUpper(s string) (result bool) {
	return strings.ToUpper(s) == s
}

func (strUtil *StrUtil) IsLower(s string) (result bool) {
	return strings.ToLower(s) == s
}

func (strUtil *StrUtil) ToPlural(s string) (result string) {
	return strUtil.pluralize.Plural(s)
}

func (strUtil *StrUtil) ToSingular(s string) (result string) {
	return strUtil.pluralize.Singular(s)
}

func (strUtil *StrUtil) ToUpper(s string) (result string) {
	return strings.ToUpper(s)
}

func (strUtil *StrUtil) ToLower(s string) (result string) {
	return strings.ToLower(s)
}

func (strUtil *StrUtil) ToCamel(s string) (result string) {
	return StrToCamel(s)
}

func (strUtil *StrUtil) ToPascal(s string) (result string) {
	return StrTitle(strUtil.ToCamel(s))
}

func (strUtil *StrUtil) ToUpperSnake(s string) (result string) {
	return strUtil.ToUpper(strUtil.ToSnake(s))
}

func (strUtil *StrUtil) ToTitle(s string) (result string) {
	return StrTitle(strUtil.splitByCapital(s, " "))
}

func (strUtil *StrUtil) ToSnake(s string) (result string) {
	return strUtil.splitByCapital(s, "_")
}

func (strUtil *StrUtil) ToKebab(s string) (result string) {
	return strUtil.splitByCapital(s, "-")
}

func (strUtil *StrUtil) splitByCapital(s, separator string) (result string) {
	pascal := strUtil.ToPascal(s)
	result = ""
	for index, ch := range pascal {
		strCh := string(ch)
		if index == 0 {
			result += strings.ToLower(strCh)
			continue
		}
		if strUtil.IsReallyUpper(strCh) {
			// only add separator for non consecutive capitals
			previousStrCh := string(pascal[index-1])
			if strUtil.IsReallyLower(previousStrCh) {
				result += separator + strings.ToLower(strCh)
				continue
			} else if index < len(pascal)-1 {
				nextStrCh := string(pascal[index+1])
				if strUtil.IsReallyLower(nextStrCh) {
					result += separator + strings.ToLower(strCh)
					continue
				}
			}
		}
		result += strings.ToLower(strCh)
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

func (strUtil *StrUtil) CurrentTime() (currentTimeStr string) {
	now := time.Now()
	return now.Format("20060102150405")
}

func (strUtil *StrUtil) EncodeBase64(str string) (result string) {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (strUtil *StrUtil) DecodeBase64(str string) (result string, err error) {
	sd, err := base64.StdEncoding.DecodeString(str)
	return string(sd), err
}

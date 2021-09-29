package utility

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/google/uuid"
)

type StrUtil struct {
	Util *Util
}

func NewStrUtil(utl *Util) *StrUtil {
	return &StrUtil{
		Util: utl,
	}
}

func (strUtil *StrUtil) IsUpper(s string) (result bool) {
	return strings.ToUpper(s) == s
}

func (strUtil *StrUtil) IsLower(s string) (result bool) {
	return strings.ToLower(s) == s
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

func (strUtil *StrUtil) EscapeShellArg(s string) (result string) {
	quoteEscapedStr := strings.ReplaceAll(s, "'", "\\'")
	return fmt.Sprintf("'%s'", quoteEscapedStr)
}

// GetSubKeys get sub keys from dictionary
func (strUtil *StrUtil) GetSubKeys(keys []string, parentKeys []string) (subKeys []string) {
	seen := map[string]bool{}
	parentKey := strings.Join(parentKeys, "::")
	prefixLength := len(parentKey) + len("::")
	subKeys = []string{}
	for _, key := range keys {
		if !strings.HasPrefix(key, parentKey+"::") {
			continue
		}
		childKey := key[prefixLength:]
		if childKey == "" {
			continue
		}
		childKeyParts := strings.SplitN(childKey, "::", 2)
		subkey := childKeyParts[0]
		seen[subkey] = true
	}
	for key := range seen {
		subKeys = append(subKeys, key)
	}
	return subKeys
}

func (strUtil *StrUtil) GetUniqueElements(arr []string) (result []string) {
	result = []string{}
	seen := map[string]bool{}
	for _, element := range arr {
		if _, exist := seen[element]; exist {
			continue
		}
		result = append(result, element)
		seen[element] = true
	}
	return result
}

// indent second-last lines
func (strUtil *StrUtil) indent(multiLineStr string, indentation string, skipFirstLine bool) (indentedStr string) {
	lines := strings.Split(multiLineStr, "\n")
	for index, line := range lines {
		if index == 0 && skipFirstLine {
			continue
		}
		if strings.Trim(line, " ") != "" {
			lines[index] = indentation + line
		}
	}
	return strings.Join(lines, "\n")
}

// indent second-last lines
func (strUtil *StrUtil) Indent(multiLineStr string, indentation string) (indentedStr string) {
	return strUtil.indent(multiLineStr, indentation, true)
}

func (strUtil *StrUtil) FullIndent(multiLineStr string, indentation string) (indentedStr string) {
	return strUtil.indent(multiLineStr, indentation, false)
}

func (strUtil *StrUtil) Replace(s string, replacementMap map[string]string) (result string) {
	result = s
	keys := []string{}
	for key := range replacementMap {
		keys = append(keys, key)
	}
	sort.Sort(ByLenDesc(keys))
	for _, key := range keys {
		re := regexp.MustCompile(key)
		val := replacementMap[key]
		result = re.ReplaceAllStringFunc(result, func(text string) string {
			indentation, _ := strUtil.GetIndentation(text, 1)
			indentedVal := strUtil.Indent(val, indentation)
			return re.ReplaceAllString(text, indentedVal)
		})
	}
	return result
}

func (strUtil *StrUtil) Repeat(s string, repetition int) (result string) {
	result = ""
	for i := 0; i < repetition; i++ {
		result += s
	}
	return result
}

func (strUtil *StrUtil) GetIndentation(s string, level int) (result string, err error) {
	rex := regexp.MustCompile("^([ \t]+).*$")
	match := rex.FindStringSubmatch(s)
	if len(match) < 2 {
		return "", fmt.Errorf("string is not predeeded by space or tab: '%s'", s)
	}
	totalIndentation := match[1]
	indentationLength := len(totalIndentation) / level
	result = s[:indentationLength]
	if strUtil.Repeat(result, level) != totalIndentation {
		return result, fmt.Errorf("cannot determine single %d indentation for '%s'", level, s)
	}
	return result, nil
}

func (strUtil *StrUtil) GetLineSubmatch(lines, patterns []string) (matchIndex int, submatch []string, err error) {
	patternIndex := 0
	rex, err := regexp.Compile(patterns[patternIndex])
	if err != nil {
		return -1, submatch, err
	}
	for lineIndex, line := range lines {
		match := rex.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		if patternIndex == len(patterns)-1 {
			return lineIndex, match, nil
		}
		patternIndex++
		rex, err = regexp.Compile(patterns[patternIndex])
		if err != nil {
			return -1, submatch, err
		}
	}
	return -1, submatch, nil
}

func (strUtil *StrUtil) prepareLinesForReplacement(lines []string) (preparedLines []string) {
	if len(lines) == 0 {
		return []string{""}
	}
	return lines
}

func (strUtil *StrUtil) ReplaceLineAtIndex(lines []string, index int, replacements []string) (result []string, err error) {
	lines = strUtil.prepareLinesForReplacement(lines)
	if index < 0 || index >= len(lines) {
		return []string{}, fmt.Errorf("index out of bound: %d", index)
	}
	tmpLines := []string{}
	tmpLines = append(tmpLines, lines[:index]...)
	tmpLines = append(tmpLines, replacements...)
	if index < len(lines) {
		tmpLines = append(tmpLines, lines[index+1:]...)
	}
	content := strings.Join(tmpLines, "\n")
	result = strings.Split(content, "\n")
	return result, nil
}

func (strUtil *StrUtil) CompleteLines(lines, patterns, suplements []string) (newLines []string, err error) {
	if len(patterns) != len(suplements) {
		return newLines, fmt.Errorf("patterns and suplements length doesn't match")
	}
	for index, pattern := range patterns {
		suplement := suplements[index]
		match, err := regexp.MatchString(pattern, suplement)
		if err != nil {
			return newLines, err
		}
		if !match {
			return newLines, fmt.Errorf("pattern[%d], %s doesn't match %s", index, pattern, suplement)
		}
	}
	newLines = append([]string{}, lines...)
	lastMatchIndex := len(newLines) - 1
	for index, suplement := range suplements {
		matchIndex, _, _ := strUtil.GetLineSubmatch(newLines, patterns[:index+1])
		if matchIndex > -1 {
			lastMatchIndex = matchIndex
			continue
		}
		newLines, _ = strUtil.ReplaceLineAtIndex(newLines, lastMatchIndex, []string{newLines[lastMatchIndex], suplement})
		lastMatchIndex, _, _ = strUtil.GetLineSubmatch(newLines, patterns[:index+1])
	}
	content := strings.Join(newLines, "\n")
	newLines = strings.Split(content, "\n")
	return newLines, nil
}

func (strUtil *StrUtil) Submatch(s string, pattern string) (result []string, err error) {
	rex, err := regexp.Compile(pattern)
	if err != nil {
		return result, err
	}
	result = rex.FindStringSubmatch(s)
	return result, err
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

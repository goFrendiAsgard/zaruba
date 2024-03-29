package strutil

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func StrQuote(s string, quote byte) (result string) {
	if len(s) > 0 && s[0] == quote && s[len(s)-1] == quote {
		return s
	}
	quoteEscapedStr := strings.ReplaceAll(s, string(quote), "\\"+string(quote))
	return fmt.Sprintf("%s%s%s", string(quote), quoteEscapedStr, string(quote))
}

func StrDoubleQuote(s string) (result string) {
	return StrQuote(s, '"')
}

func StrSingleQuote(s string) (result string) {
	return StrQuote(s, '\'')
}

func StrEscapeShellValue(value string) (result string) {
	escapedValue := strings.ReplaceAll(value, "'", "'\"'\"'")
	return fmt.Sprintf("'%s'", escapedValue)
}

func StrShellVariable(key, value string) (result string) {
	return fmt.Sprintf("%s=%s", key, StrEscapeShellValue(value))
}

func strIndent(multiLineStr string, indentation string, skipFirstLine bool) (indentedStr string) {
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
func StrIndent(multiLineStr string, indentation string) (indentedStr string) {
	return strIndent(multiLineStr, indentation, true)
}

// indent all lines
func StrFullIndent(multiLineStr string, indentation string) (indentedStr string) {
	return strIndent(multiLineStr, indentation, false)
}

func StrRepeat(s string, repetition int) (result string) {
	result = ""
	for i := 0; i < repetition; i++ {
		result += s
	}
	return result
}

func StrGetIndentation(s string, level int) (result string, err error) {
	rex := regexp.MustCompile("^([ \t]+).*$")
	match := rex.FindStringSubmatch(s)
	if len(match) < 2 {
		return "", nil
	}
	totalIndentation := match[1]
	indentationLength := len(totalIndentation) / level
	result = s[:indentationLength]
	if StrRepeat(result, level) != totalIndentation {
		return result, fmt.Errorf("cannot determine single %d indentation for '%s'", level, s)
	}
	return result, nil
}

func StrReplace(s string, replacementMap map[string]string) (result string) {
	result = s
	keys := []string{}
	for key := range replacementMap {
		keys = append(keys, key)
	}
	sort.Sort(ReplacementMapKey(keys))
	for _, key := range keys {
		val := replacementMap[key]
		if re, err := regexp.Compile(key); err == nil {
			result = re.ReplaceAllStringFunc(result, func(text string) string {
				indentation, _ := StrGetIndentation(text, 1)
				indentedVal := StrFullIndent(val, indentation)
				return re.ReplaceAllString(text, indentedVal)
			})
		} else {
			result = strings.ReplaceAll(result, key, val)
		}
	}
	return result
}

func StrGetLineSubmatch(lines, patterns []string, desiredPatternIndex int) (matchIndex int, submatch []string, err error) {
	patternIndex := 0
	if desiredPatternIndex < 0 {
		desiredPatternIndex = len(patterns) + desiredPatternIndex
	}
	rex, err := regexp.Compile(patterns[patternIndex])
	if err != nil {
		return -1, submatch, err
	}
	for lineIndex, line := range lines {
		match := rex.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		if patternIndex == desiredPatternIndex {
			matchIndex = lineIndex
			submatch = match
		}
		if patternIndex == len(patterns)-1 {
			return matchIndex, submatch, nil
		}
		patternIndex++
		rex, err = regexp.Compile(patterns[patternIndex])
		if err != nil {
			return -1, submatch, err
		}
	}
	return -1, submatch, nil
}

func strPrepareLinesForReplacement(lines []string) (preparedLines []string) {
	if len(lines) == 0 {
		return []string{""}
	}
	return lines
}

func StrReplaceLineAtIndex(lines []string, index int, replacements []string) (result []string, err error) {
	lines = strPrepareLinesForReplacement(lines)
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

func StrSubmatch(s string, pattern string) (result []string, err error) {
	rex, err := regexp.Compile(pattern)
	if err != nil {
		return result, err
	}
	result = rex.FindStringSubmatch(s)
	return result, err
}

func StrTitle(s string) string {
	// This function is
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}

func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}

func StrToCamel(s string) (result string) {
	rex := regexp.MustCompile("[^a-zA-Z0-9]+")
	strippedStr := rex.ReplaceAllString(s, " ")
	titledStr := StrTitle(strippedStr)
	result = strings.ReplaceAll(titledStr, " ", "")
	if len(result) > 0 {
		firstLetter := strings.ToLower(string(result[0]))
		rest := result[1:]
		return firstLetter + rest
	}
	return result
}

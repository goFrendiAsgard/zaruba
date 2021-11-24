package strutil

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
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

func StrEscapeShellArg(s string) (result string) {
	return StrSingleQuote(s)
}

func StrShellVariable(key, value string) (result string) {
	return fmt.Sprintf("%s=%s", key, StrSingleQuote(value))
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
	sort.Sort(ByLenDescending(keys))
	for _, key := range keys {
		re := regexp.MustCompile(key)
		val := replacementMap[key]
		result = re.ReplaceAllStringFunc(result, func(text string) string {
			indentation, _ := StrGetIndentation(text, 1)
			indentedVal := StrFullIndent(val, indentation)
			return re.ReplaceAllString(text, indentedVal)
		})
	}
	return result
}

func StrGetLineSubmatch(lines, patterns []string) (matchIndex int, submatch []string, err error) {
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

func StrCompleteLines(lines, patterns, suplements []string) (newLines []string, err error) {
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
		matchIndex, _, _ := StrGetLineSubmatch(newLines, patterns[:index+1])
		if matchIndex > -1 {
			lastMatchIndex = matchIndex
			continue
		}
		newLines, _ = StrReplaceLineAtIndex(newLines, lastMatchIndex, []string{newLines[lastMatchIndex], suplement})
		lastMatchIndex, _, _ = StrGetLineSubmatch(newLines, patterns[:index+1])
	}
	content := strings.Join(newLines, "\n")
	newLines = strings.Split(content, "\n")
	return newLines, nil
}

func StrSubmatch(s string, pattern string) (result []string, err error) {
	rex, err := regexp.Compile(pattern)
	if err != nil {
		return result, err
	}
	result = rex.FindStringSubmatch(s)
	return result, err
}

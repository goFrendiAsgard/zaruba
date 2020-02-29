package file

import (
	"path"

	"github.com/zealic/xignore"
)

// GetAllFiles fetch sub-files and sub-directories recursively
func GetAllFiles(dirName string, option *Option) (allFiles []string, err error) {
	allFiles = []string{dirName}
	result, err := xignore.DirMatches(dirName, &xignore.MatchesOptions{
		Ignorefile: ".gitignore",
		Nested:     true, // Handle nested ignorefile
	})
	if err != nil {
		return allFiles, err
	}
	// create xignorePatterns
	xignorePatterns := []*xignore.Pattern{}
	for _, strPattern := range option.GetIgnores() {
		xignorePattern := xignore.NewPattern(strPattern)
		if err = xignorePattern.Prepare(); err != nil {
			return allFiles, err
		}
		xignorePattern.Prepare()
	}
	// add all sub-directories that doesn't match gitignore
	for _, subDirName := range result.UnmatchedDirs {
		absSubDirName := path.Join(dirName, subDirName)
		match := false
		for _, xignorePattern := range xignorePatterns {
			if xignorePattern.Match(absSubDirName) {
				match = true
				break
			}
		}
		if !match {
			allFiles = append(allFiles, absSubDirName)
		}
	}
	if !option.GetIsOnlyDir() {
		// add all sub-files that doesn't match gitignore
		for _, subFileName := range result.UnmatchedFiles {
			absFileName := path.Join(dirName, subFileName)
			match := false
			for _, xignorePattern := range xignorePatterns {
				if xignorePattern.Match(absFileName) {
					match = true
					break
				}
			}
			if !match {
				allFiles = append(allFiles, absFileName)
			}
		}
	}
	return allFiles, err
}

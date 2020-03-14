package file

import (
	"path"

	"github.com/zealic/xignore"
)

// GetAllFiles fetch sub-files and sub-directories recursively
func GetAllFiles(dirName string, option *Option) (allFiles []string, err error) {
	allFiles = []string{dirName}
	// create pattern
	result, err := xignore.DirMatches(dirName, &xignore.MatchesOptions{
		Ignorefile: ".gitignore",
		Nested:     true, // Handle nested ignorefile
	})
	if err != nil {
		return allFiles, err
	}
	// load ignoreConfig
	i := NewIgnoreConfig(dirName)
	// add all sub-directories that doesn't match gitignore and not in IgnoreConfig
	for _, subDirName := range result.UnmatchedDirs {
		absSubDirName := path.Join(dirName, subDirName)
		if !i.ShouldIgnore(absSubDirName) {
			allFiles = append(allFiles, absSubDirName)
		}
	}
	if !option.GetIsOnlyDir() {
		// add all sub-files that doesn't match gitignore and not in IgnoreConfig
		for _, subFileName := range result.UnmatchedFiles {
			absFileName := path.Join(dirName, subFileName)
			if !i.ShouldIgnore(absFileName) {
				allFiles = append(allFiles, absFileName)
			}
		}
	}
	return allFiles, err
}

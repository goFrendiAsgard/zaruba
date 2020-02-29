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
	// add all sub-directories that doesn't match gitignore
	for _, subDirName := range result.UnmatchedDirs {
		allFiles = append(allFiles, path.Join(dirName, subDirName))
	}
	if !option.GetIsOnlyDir() {
		// add all sub-files that doesn't match gitignore
		for _, subFileName := range result.UnmatchedFiles {
			allFiles = append(allFiles, path.Join(dirName, subFileName))
		}
	}
	return allFiles, err
}

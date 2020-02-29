package file

import (
	"log"
	"path"
	"strings"

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
	ignores := option.GetIgnores()
	// add all sub-directories that doesn't match gitignore
	for _, subDirName := range result.UnmatchedDirs {
		absSubDirName := path.Join(dirName, subDirName)
		ignored := false
		for _, prefix := range ignores {
			if strings.HasPrefix(absSubDirName, prefix) {
				ignored = true
				log.Println(absSubDirName, prefix, ignored)
				break
			}
		}
		log.Println(absSubDirName, ignored)
		if !ignored {
			log.Println(absSubDirName)
			allFiles = append(allFiles, absSubDirName)
		}
	}
	if !option.GetIsOnlyDir() {
		// add all sub-files that doesn't match gitignore
		for _, subFileName := range result.UnmatchedFiles {
			absFileName := path.Join(dirName, subFileName)
			ignored := false
			for _, prefix := range ignores {
				if strings.HasPrefix(absFileName, prefix) {
					ignored = true
					break
				}
			}
			if !ignored {
				allFiles = append(allFiles, absFileName)
			}
		}
	}
	return allFiles, err
}

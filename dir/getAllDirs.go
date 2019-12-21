package dir

// GetAllDirs fetch sub directories recursively
func GetAllDirs(dirPath string) (allDirs []string, err error) {
	return GetAllFiles(
		dirPath,
		GetFilesOption{
			MaxDepth: DefaultGetFilesOption.MaxDepth,
			OnlyDir:  true,
		},
	)
}

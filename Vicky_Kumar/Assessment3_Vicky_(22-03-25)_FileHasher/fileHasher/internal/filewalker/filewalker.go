package filewalker

import (
	"os"

	"fileHasher/internal/utils"
)

type FileWalker struct{}

func NewFileWalker() *FileWalker {
	return &FileWalker{}
}

func (fw *FileWalker) ReadDirectory(dirPath string) ([]string, error) {
	if dirPath == "" {
		return nil, utils.NewCustomError(utils.ErrEmptyDirPath)
	}

	fileInfo, err := os.ReadDir(dirPath)

	if err != nil {
		return nil, utils.NewCustomError(utils.ErrReadingDir + err.Error())
	}

	var files []string
	for _, file := range fileInfo {
		files = append(files, dirPath+"/"+file.Name())
	}

	return files, nil
}

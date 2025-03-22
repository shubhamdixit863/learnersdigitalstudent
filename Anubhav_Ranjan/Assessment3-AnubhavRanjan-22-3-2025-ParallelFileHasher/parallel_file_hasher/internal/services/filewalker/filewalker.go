package filewalker

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	errInvalidDirPath = "provided directory path is not a directory"
	errReadDir        = "error reading directory: %w"
)

type FileWalker struct {
	dirPath string
}

func NewFileWalker(dirPath string) *FileWalker {
	return &FileWalker{
		dirPath: dirPath,
	}
}

func (fw *FileWalker) Walk() ([]string, error) {
	fileInfo, err := os.Stat(fw.dirPath)
	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		return nil, fmt.Errorf(errInvalidDirPath)
	}

	var files []string
	err = filepath.Walk(fw.dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf(errReadDir, err)
	}

	return files, nil
}

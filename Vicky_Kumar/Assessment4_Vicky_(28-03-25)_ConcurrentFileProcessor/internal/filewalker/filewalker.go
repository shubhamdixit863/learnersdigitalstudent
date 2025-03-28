package filewalker

import (
	"concurrent_file_processor/internal/utils"
	"log"
	"os"
	"strings"
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
func (fw *FileWalker) ReadFileLines(filePath string, linesCh chan string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("%s (%s) : %v \n", utils.ErrReadingFile, filePath, err)
		return
	}

	lines := strings.Split(string(content), "\n")
	// fan in
	for _, line := range lines {
		linesCh <- line
	}

}

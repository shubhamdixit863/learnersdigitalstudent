package hasher

import (
	"crypto/sha256"
	"os"

	"fileHasher/internal/utils"
)

type FileHasher struct{}

func NewFileHasher() *FileHasher {
	return &FileHasher{}
}

func (fh *FileHasher) HashFile(filePath string) (string, error) {
	if filePath == "" {
		return "", utils.NewCustomError(utils.ErrEmptyFilePath)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", utils.NewCustomError(utils.ErrReadingFile + err.Error())
	}

	h := sha256.New()
	h.Write(content)
	return string(h.Sum(nil)), nil
}

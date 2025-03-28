package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func FindTextFiles(dir string) ([]string, error) {
	var txtFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".txt") {
			txtFiles = append(txtFiles, path)
		}

		return nil
	})

	return txtFiles, err
}

func CountWords(line string) int {
	return len(strings.Fields(line))
}

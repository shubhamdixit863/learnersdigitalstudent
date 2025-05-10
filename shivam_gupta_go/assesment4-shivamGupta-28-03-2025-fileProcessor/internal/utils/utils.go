package utils

import (
	"os"
	"path/filepath"
)
const (
	txt =".txt"
)


func GetTxtFiles(dir string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) ==  txt{
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}

	return files, nil
}
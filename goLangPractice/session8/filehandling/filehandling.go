package filehandling

import (
	"os"
)

func Readfile(filename string) (string, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

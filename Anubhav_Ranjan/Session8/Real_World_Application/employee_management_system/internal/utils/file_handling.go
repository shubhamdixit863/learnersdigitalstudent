package utils

import (
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, os.ModePerm)
}

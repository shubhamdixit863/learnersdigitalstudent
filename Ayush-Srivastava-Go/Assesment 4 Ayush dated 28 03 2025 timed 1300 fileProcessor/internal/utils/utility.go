package utils

import (
	"path/filepath"
)

func SplitLines(data string) []string {
	return filepath.SplitList(data) 
}
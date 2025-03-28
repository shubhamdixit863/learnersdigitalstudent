package services

import (
	"strings"
)


func Chunker(line string) []string {
	return strings.Fields(line) 
}

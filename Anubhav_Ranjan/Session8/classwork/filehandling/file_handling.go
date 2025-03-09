package filehandling

import "os"

func Readfile(filname string) (string, error) {
	data, err := os.ReadFile(filname)
	if err != nil {
		return "", err
	}
	return string(data), err
}

package services

import (
	
	"io/fs"
	"math/rand"
	
	"strings"
	"time"
	"path/filepath"
)

func GetTxtFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".txt") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func MakeAPICall(line string) string {
	
	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(200)))
	return "200 OK"
}

func RetryableAPICall(line string, retries int) string {
	for i := 0; i < retries; i++ {
		status := MakeAPICall(line)
		if strings.HasPrefix(status, "200") {
			return status
		}
		time.Sleep(time.Second * time.Duration(i+1)) 
	}
	return "Failed after retries"
}



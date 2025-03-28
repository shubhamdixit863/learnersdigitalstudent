package filehandler

import (
	"fileprocessor/internal/utils"
	"io/fs"
	"os"
	"path/filepath"
)

//get all the `.txt` files from the directory
func GetFiles(dir string) ([]string, error){

	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == utils.SearchedFileFormat {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

//Read the contents
func ReadFile(path string) ([]string, error){
	
	data, err :=  os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	lines := utils.SplitLines(string(data))

	return lines, nil
}



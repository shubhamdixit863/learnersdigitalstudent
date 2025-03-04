package filehandling

import (
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(data), err

}

// func CreateDirectiry(filename string) (string, error) {
// 	os.Mkdir(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return data, err

// }

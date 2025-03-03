package filer

//readfile

import (
	//"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(file))

}
func WriteFile(filename string, data []byte) {
	err := os.WriteFile(filename, data, 0666)
	if err != nil {
	}

}

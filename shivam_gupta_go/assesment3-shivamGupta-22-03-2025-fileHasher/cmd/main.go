package main

import (
	"fmt"
	"log"

	"fileHasher/internal/filewalker"
	"fileHasher/internal/services"
	"fileHasher/internal/utility"
)

func main() {
    
   errmsg1:=utility.Errmsg1
	 dir:=utility.Dir
	 printRes:=utility.PrintRes
    
	 
    files, err := filewalker.WalkDirectory(dir)
    if err != nil {
        log.Fatalf(errmsg1, err)
    }

    
    results := services.HashFiles(files)

    
    for _, result := range results {
        fmt.Printf(printRes, result.File, result.Hash)
    }
}
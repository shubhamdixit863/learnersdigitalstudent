package config

import (
	"concurrent-file-processor/constants"
	"flag"
	"fmt"
	"os"
)

// Arguments struct
type Arguments struct {
	DirPath string
	Mode    string
	Keyword string
}

// ParseArgs handles command-line argument parsing
func ParseArgs() Arguments {
	dirPath := flag.String("dir", "", "Path to directory containing text files")
	mode := flag.String("mode", "filter", "Processing mode: filter | wordcount | apicall | retryapi")
	keyword := flag.String("keyword", "error", "Keyword for filter mode")
	flag.Parse()

	if *dirPath == "" {
		fmt.Println(constants.MissingDirPath)
		os.Exit(1)
	}

	return Arguments{DirPath: *dirPath, Mode: *mode, Keyword: *keyword}
}

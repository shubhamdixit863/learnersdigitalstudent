package utils

import (
	"concurrent_file_processing/internals/models"
	"os"
)

const SplitWord = "error"
const ApiPath = "https://httpbin.org/post"

const DirName = "text"

const NumberOfWorker = 3

var FileInput = make(chan os.DirEntry, NumberOfWorker)

var FileDataOutput = make(chan *models.ProcessedData, NumberOfWorker)
var (
	Outputmessage = "you Directory of Text Has been processed \n enter ypur choice for foolowing type of output\n" +
		"1. Line Filter " +
		"2. Word Count" +
		"3. Api Call" +
		"4. Retryable Api Call"
)

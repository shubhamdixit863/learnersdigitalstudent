package utils

import (
	"concurrent_file_processing/internals/models"
	"os"
)

const SplitWord = "error"
const ApiPath = "https://httpbin.org/post"

const NumberOfWorker = 3

var FileInput = make(chan os.DirEntry, NumberOfWorker)

var FileDataOutput = make(chan *models.ProcessedData, NumberOfWorker)

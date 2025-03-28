package services

import (
	"fmt"
	"log"
	"omagrawalAssesment4-28-03-25/internals/models"
	"omagrawalAssesment4-28-03-25/internals/utils"
	"strings"
	"sync"
)

var (
	WorkerId = "currently executing worker of id"
)

func CreatWorkerPool(wg *sync.WaitGroup) {
	id := 0
	for range utils.NumberOfWorker {
		wg.Add(1)
		go WorkerFunc(id, wg)
		id++
	}
}

func WorkerFunc(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println(WorkerId, id)
	for file := range utils.FileInput {
		filedata := ReadFileFunc(file)
		line := strings.Split(filedata, "\n")
		for n, linedata := range line {
			utils.FileDataOutput <- models.NewProcessedData(file.Name(), linedata, n)

		}
	}
}

func ProcessAllData() {
	fmt.Println(utils.Outputmessage)
	var option int
	fmt.Scanln(&option)

	for textLine := range utils.FileDataOutput {
		switch option {
		case 1:
			LineFilter(textLine)
		case 2:
			WordCount(textLine)
		case 3:
			PostToApi(textLine)
		case 4:
			RetryPostToApi(textLine)

		}

	}
}

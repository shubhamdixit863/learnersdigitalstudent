package services

import (
	"concurrent_file_processing/internals/models"
	"concurrent_file_processing/internals/utils"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
)

var (
	WorkerId = "currently executing worker of id"
)
var arr = []string{"Line Filter,", "Word Count,", "Api Call,", "Retryable Api Call"}

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
	RandomIntegerwithinRange := rand.Intn(3)
	var option int
	option = RandomIntegerwithinRange
	fmt.Println("operatin performed is ", option+1, arr)
	fmt.Println(arr)

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

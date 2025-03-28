package processor

import (
	"fileprocessor/internal/filehandler"
	"fileprocessor/internal/services"
	"fileprocessor/internal/utils"
	"log"
	"sync"
)

//Process each file with a new go routine
func ProcessFiles(dir string, mode string) error {

	files, err := filehandler.GetFiles(dir)

	if err != nil {
		return err
	}

	lineChannel := make(chan string)
	var wg sync.WaitGroup

	// Process each file with a new go routine(fan-out)
	for _, file := range files {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			lines, err := filehandler.ReadFile(filePath)
			if err != nil {
				log.Println(utils.FileReadingError, err)
				return
			}
			for _, line := range lines {
				lineChannel <- line
			}
		}(file)
	}

	go func(){
		wg.Wait()
		close(lineChannel)
	}()

	return services.ProcessLines(lineChannel, mode)

}
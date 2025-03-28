package services

import (
	"fileprocessor/internal/utils"
	"fmt"
	"strings"
	"sync"
)

//ProcessLines processes each line
func ProcessLines(lineChannel <-chan string, mode string) error {

	switch mode {
	case utils.Modes[0]:
		CountWords(lineChannel)
	case utils.Modes[1]:
		FilterLines(lineChannel, utils.FilteredText)
	case utils.Modes[2]:
		APICall(lineChannel)
	default:
		fmt.Println(utils.InvalidModeError)
	}
	return nil
}

//CountWords counts the number of words in each line
func CountWords(lineChannel <-chan string){

	var count int
	for line := range lineChannel {
		count += len(strings.Fields(line))
	}	

	fmt.Println(utils.TotalWords, count)
}


//FilterLines filters the lines based on a conditioned string
func FilterLines(lineChannel <-chan string, keyword string){

	for line := range lineChannel {
		if strings.Contains(line, keyword) {
			fmt.Println(utils.Filtered, line)
		}
	}

}

//APICall makes an API call for each line
func APICall(lineChannel <-chan string){
	var wg sync.WaitGroup

	for line := range lineChannel{
		wg.Add(1)
		go func(data string){
			defer wg.Done()

			status := utils.SendAPI(data, 3)
			fmt.Println(utils.APIStatus, status)

		}(line)
	}
	wg.Wait()
}
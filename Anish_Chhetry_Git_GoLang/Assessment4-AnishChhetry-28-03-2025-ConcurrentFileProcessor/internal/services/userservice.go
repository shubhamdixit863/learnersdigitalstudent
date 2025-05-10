package services

import "fmt"

const (
	spaceString  = " "
	emptyString  = ""
	inFileString = "in file:"
	int0         = 0
	int3         = 3
	optionAsk    = `What do you want to do?
1. Word Count
2. Line Filer
3. Call API
4. Exit`
	filterWord = "error"
	api        = "https://httpbin.org/post"
)

func User(storage map[string]string) {
	var option int
	for {
		fmt.Println(optionAsk)

		fmt.Scanln(&option)

		switch option {
		case 1:
			WordCount(storage)
		case 2:
			LineFilter(storage, filterWord)
		case 3:
			RetryCallAPI(storage, api)
		case 4:
			return
		}

	}
}

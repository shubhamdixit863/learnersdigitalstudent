package main

import (
	"fmt"
	"session-revision/get"
	"session-revision/put"
)

func main() {
	for {
		var choose int
		fmt.Println("Press 1 to put 2 to get 3 to exit")
		fmt.Scanln(&choose)
		switch choose {
		case 1:
			var key int
			fmt.Println("enter key")
			fmt.Scanln(&key)
			var val string
			fmt.Println("enter value")
			fmt.Scanln(&val)
			put.Putting(key, val)
		case 2:
			var key int
			fmt.Println("enter key")
			fmt.Scanln(&key)
			get.Getting(key)
		case 3:
			return
		}

	}
}

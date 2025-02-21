package start

import (
	"fmt"
)

var Answer int

func Hello() {
	for {

		var count = 0
		fmt.Println("1. What is the purpose of the go.mod file in a Go project?")
		fmt.Println("a) To store environment variables \t")
		fmt.Println("b) To manage dependencies and module information ")
		fmt.Println("c) To define functions and packages\t ")
		fmt.Println("d) To compile the Go code ")
		var ans string
		fmt.Println("choice frome above options")
		fmt.Scan(&ans)

		if ans == "b" {
			count++
		}

		fmt.Println("1. What is the purpose of the go.mod file in a Go project?")
		fmt.Println("a) To store environment variables \t")
		fmt.Println("b) To manage dependencies and module information ")
		fmt.Println("c) To define functions and packages\t ")
		fmt.Println("d) To compile the Go code ")

		fmt.Println("choice frome above options")
		fmt.Scan(&ans)

		if ans == "b" {
			count++
		}
		fmt.Println("1. What is the purpose of the go.mod file in a Go project?")
		fmt.Println("a) To store environment variables \t")
		fmt.Println("b) To manage dependencies and module information ")
		fmt.Println("c) To define functions and packages\t ")
		fmt.Println("d) To compile the Go code ")

		fmt.Println("choice frome above options")
		fmt.Scan(&ans)

		if ans == "b" {
			count++
		}
		fmt.Println("1. What is the purpose of the go.mod file in a Go project?")
		fmt.Println("a) To store environment variables \t")
		fmt.Println("b) To manage dependencies and module information ")
		fmt.Println("c) To define functions and packages\t ")
		fmt.Println("d) To compile the Go code ")

		fmt.Println("choice frome above options")
		fmt.Scan(&ans)

		if ans == "b" {
			count++
		}
		fmt.Println("1. What is the purpose of the go.mod file in a Go project?")
		fmt.Println("a) To store environment variables \t")
		fmt.Println("b) To manage dependencies and module information ")
		fmt.Println("c) To define functions and packages\t ")
		fmt.Println("d) To compile the Go code ")

		fmt.Println("choice frome above options")
		fmt.Scan(&ans)

		if ans == "b" {
			count++
		}

		fmt.Println(count)

		Answer = count
	}

}

package score

import "fmt"

var marks int = 0

func ViewScore() {
	fmt.Printf("You got %d marks out of 5.", &marks)
}

func IncrementScore() {
	marks++
}
func ResetScore() {
	marks = 0
}

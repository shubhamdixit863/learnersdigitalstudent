package main

import (
	"Quiz_Game/display"
	"Quiz_Game/start"
	"fmt"
)
func main()  {
	var answer int
	for{
	fmt.Println(display.Display())
	fmt.Scan(&answer)
	if answer == 1 {
		start.Start()
	}else if answer == 2 {
		fmt.Println(start.Score)
	}else if answer == 3 {
		break
	}
}
}
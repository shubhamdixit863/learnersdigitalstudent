package exit

import (
	"fmt"
	"module_Assignment/score"
)

var exit =false

func Exit() {
	score.Score=0;
	exit=true
	fmt.Println("-----SuccessFully Exited the game -----")

}

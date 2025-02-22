package start

import (
	"Quiz_Game/answer"
	"Quiz_Game/question"
	"fmt"
)
var Score int
func Start() {
	var given_answer int
	for i:=0;i<len(question.Question());i++{
		fmt.Println(question.Question()[i])
		
		fmt.Scan(&given_answer)
		if given_answer == answer.Answer()[i]{
			Score++
		}
	}
	
	fmt.Println("Final Score:", Score)
}
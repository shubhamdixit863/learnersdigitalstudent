package quiz
import "fmt"

var Data=0
func StartQuiz()int{
	var answer=[2]int{3,2}
	var ans int
	var score int=0
	fmt.Println("Question 1: What is the capital of France?")
	fmt.Println("Berlin");
	fmt.Println("Madrid");
	fmt.Println("Paris");
	fmt.Println("Rome");
	fmt.Scanln(&ans)
	if(ans==answer[0]){
		score+=1
	}
	fmt.Println("Question 2: 2+2?")
	fmt.Println("44");
	fmt.Println("4");
	fmt.Println("6");
	fmt.Println("2");
	fmt.Scanln(&ans)
	if(ans==answer[1]){
		score+=1
	}
	return score



}
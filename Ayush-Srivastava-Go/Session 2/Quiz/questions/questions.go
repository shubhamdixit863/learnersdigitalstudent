package questions

import "fmt"

var Questions = []string{
	"What is the capital of France?",
	" 2 + 2 = ?",
	"What is the esports game of the year?",
	"How many times did messi win ballon d'or?",
	"What is going on?",
}

var Options = [][]string{
	{"Berlin", "Madrid", "Paris", "Rome"},
	{"1", "2", "3", "4"},
	{"CS2", "Valorant", "Pokemon Go", "PUBG"},
	{"7", "8", "9", "10"},
	{"Who cares", "I dunno", "Why asking", "Its all fked up"},
}

var answers = []int{2, 3, 1, 1, 3}
var Score int

func StartQuiz() {
	Score = 0

	for i := 0; i < 5; i++ {
		fmt.Printf("\nQuestion %d: %s\n", i+1, Questions[i])
		for j:=0; j<4; j++{
			fmt.Printf("%d) %s\n", j+1, Options[i][j])
		}
		fmt.Printf("Your answer: ")
		userAnswer := Validate(1, 2, 3, 4)
		if(isCorrect(answers[i], userAnswer-1)){
			fmt.Println("Correct!")
			Score++
		}else{
			fmt.Printf("Oops! Wrong Answer. The correct answer was %s\n", Options[i][answers[i]])
		}
	}

	fmt.Printf("\n** Quiz Completed! **\n You scored %d out 5\n", Score)
}

func Validate(allowedOptions ...int) int {
	var input int

	for{
		_, err := fmt.Scanln(&input)
		if err == nil {
			for _, option := range allowedOptions{
				if input == option{
					return input
				}
			}
		}

		fmt.Println("Invlid Choice!")
	}
}

func isCorrect(correctOption int, userChoice int) bool {
	return correctOption == userChoice
}

func ViewScore(){
	fmt.Printf("You scored %d out of 5!\n", Score)
}
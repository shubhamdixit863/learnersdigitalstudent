package quize

import (
	"fmt"
	"module_Assignment/score"
)

func Quize() {

	fmt.Println(score.Score)
	fmt.Println("enetr integer")
	fmt.Println("Question 1")
	fmt.Println("who is president of india \n 1.modi    \n 2.murmu  \n 3.yogi \n 4.amit shah")
	fmt.Println("enetrr your answer")
	var ans1 int
	fmt.Scanln(&ans1)

	if ans1 == 2 {
		score.Score += 1
		fmt.Println("correct answer")

	} else {
		fmt.Println("wrong answer")
	}

	fmt.Println("Question 2")
	fmt.Println("who is PM of india \n 1.modi    \n 2.modi \n 3.yogi \n 4.amit shah")
	fmt.Println("enetrr your answer")
	var ans2 int
	fmt.Scanln(&ans2)

	if ans2 == 1 {
		score.Score += 1
		fmt.Println("correct answer")

	} else {
		fmt.Println("wrong answer")
	}

	fmt.Println("Question 1")
	fmt.Println("who is Cm of UP \n 1.modi    \n 2.modi \n 3.yogi \n 4.amit shah")
	fmt.Println("enetrr your answer")
	var ans3 int
	fmt.Scanln(&ans3)

	if ans3 == 3 {
		score.Score += 1
		fmt.Println("correct answer")

	} else {
		fmt.Println("wrong answer")
	}
}

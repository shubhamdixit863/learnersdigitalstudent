package start 
import (
	"fmt"
	"question2/quiz"
)

func Start () {
	for {
		fmt.Println("1.start quiz")
	fmt.Println("2. view score ")
	fmt.Println("3. Exit ")
	var a int
	fmt.Scanln(&a)
	if a==1{
     quiz.Quiz()
	}else if a==2 { 

    fmt.Printf("score : %d\n", quiz.Score)
	}else if a==3 {
     break
	}else {
		fmt.Println("invalid input")
	}
	}
  
}
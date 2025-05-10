package main 
import "fmt"

func main(){
	fizzBuzz(3)
	fizzBuzz(5)
	fizzBuzz(15)
	fizzBuzz(4)
}

func fizzBuzz (num int ){

 if num%3==0 && num%5==0{
		fmt.Println("FizzBuzz")}else if  num%3==0 {
		fmt.Println("Fizz")
	}else if num%5==0 {
    fmt.Println("Buzz")
	}else {
		fmt.Println("error")
	}
	
}
package main
import ("fmt"
"os")

func main(){
  fmt. Println("Enter the name, age and bank balance")
  name := os.Args[1]
  age := os.Args[2]
  bankBalance := os.Args[3]

  fmt.Println("Name: ", name)
  fmt.Println("Age: ", age)
  fmt.Println("Bank Balance: ", bankBalance)
 
}
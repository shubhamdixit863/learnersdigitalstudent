package main 
import "fmt"



func main(){
var name string
var age int
var address string
var company string 
var parents_name string
fmt.Println("enter name" )
fmt.Scan(&name)
fmt.Println("enter age ")
fmt.Scan(&age )

fmt.Println("enter address ")
fmt.Scan( &address)

fmt.Println("enter company")
fmt.Scan( &company )

fmt.Println("enter parents name ")
fmt.Scan(&parents_name)

age = age*365;
fmt.Println("name ", name)
fmt.Println("age ", age )
fmt.Println("address " ,address)
fmt.Println("company name ", company)
fmt.Println("parents_name", parents_name)

}
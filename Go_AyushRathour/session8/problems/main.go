package main

import ("fmt")
 type Data struct{
	name string
	age int
 }


 //value receiver
 func (d Data)DoubleAge()int{
     d.age=2 *d.age
	 return d.age
 }

 //pointer receiever
 func (d *Data)DoubleAgePointer()int{
	d.age=2 *d.age
	return d.age
}


 func main(){
   data:= Data{
	name: "Jhon",
	age: 23,
   }

   c:=data.DoubleAge()
   fmt.Println(c)
   fmt.Println(data.age)

   j:=data.DoubleAgePointer()
   fmt.Println(j)
   fmt.Println(data.age)



 }
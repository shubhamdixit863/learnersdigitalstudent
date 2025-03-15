package main

import (
	"fmt"
	"sync"
)

func Mul(i int,Wg *sync.WaitGroup) {
	fmt.Println("processing : ", i,"->",i*2)
  defer Wg.Done()
}

func main() {
	var Wg sync.WaitGroup

  numbers:=[]int{1,2,3,4,5}
	
	for i:=0;i<len(numbers);i++{
		go Mul(i+1,&Wg)
	}

	Wg.Add(5)
	Wg.Wait()
	fmt.Println("All routines completed")
}

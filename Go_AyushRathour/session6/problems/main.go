package main

import "fmt"

func main(){
	// var slc []int
	// fmt.Println(slc)
	// slc=append(slc, 2)
	// fmt.Println(len(slc), cap(slc))

	// slc2:= make([]int, 2)
	// fmt.Println(slc2)
    fmt.Println("1. Slice Capacity & Append Behavior")
	question1()
	fmt.Println("2. Reslicing and Capacity Growth")
	question2()
	fmt.Println("3. Slice Reference Issue")
	question3()
	fmt.Println("4. Copying vs. Appending")
	question4()
	fmt.Println("5. Slice Capacity Overflow")
	question5()
	

}


// 1. Slice Capacity & Append Behavior

func question1(){
		arr := []int{1, 2, 3, 4, 5}
	
		s1 := arr[:2]  //makes a slice of existing arr slice
	
		s2 := arr[1:3]   //makes a slice of existing arr slice
		
		s1 = append(s1, 10) // Modify slice s1, hence it also reflects in arr  
	
		fmt.Println(arr, s1, s2) // here arr is cahnged and s2 slice is also modified
	
	
}
//2. Reslicing and Capacity Growth

func question2() {

	arr := []int{1, 2, 3, 4, 5}

	s := arr[1:3]    //slice of arr with len 2 and cap 4

	fmt.Println(len(s), cap(s)) // Length and capacity before appending

    

	s = append(s, 10, 20, 30) //new slice 

	fmt.Println(arr, s) //therefore it didnt reflect in previous arr slice
	fmt.Println(len(s), cap(s))

}

//3. Slice Reference Issue

func modifySlice(s []int) {

    s[0] = 100

}



func question3() {

    arr := []int{1, 2, 3, 4, 5}

    s := arr[:3] // reference of slice arr is taken and a slice s is created

     

    modifySlice(s)  //slice is passed as pass by reference

    

    fmt.Println(arr)

}

// 4. Copying vs. Appending
func question4() {

    s1 := []int{1, 2, 3, 4}

    s2 := make([]int, len(s1))

    copy(s2, s1) // new slice s2 is created 



    s1[0] = 100

    fmt.Println(s1, s2) //two different slices are created



    s1 = append(s1, 200)  // changes in slice 1 is done

    fmt.Println(s1, s2)  //as there are two different slice changes in one slice does not effect the other slice

}


// 5. Slice Capacity Overflow

func question5() {

    s := make([]int, 2, 3) //len 2 cap 3  s=[0,0]
    s[0], s[1] = 10, 20    // s=[10, 20]



    s1 := append(s, 30) // s1 is same slice as s so changes in this slice refelcts in s s=s1=[10,20, 30]

    s2 := append(s, 40)  //new slice is created as cap is limit is exceeded.



    fmt.Println(s, s1, s2)

}
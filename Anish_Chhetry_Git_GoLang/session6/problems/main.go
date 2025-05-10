package main

import "fmt"

func main() {
	q1()
	q2()
	q3()
	q4()
	q5()
}
func q1() {
	// 1. Slice Capacity & Append Behavior

	// Problem:

	// What will be the output of the following Go program? Explain why.

	arr := []int{1, 2, 3, 4}
	s1 := arr[:2]
	s2 := arr[1:3]
	s1 = append(s1, 10) // Modify slice s1
	fmt.Println(arr, s1, s2)

	//Output:
	//[1 2 10 4] [1 2 10] [2 10]
	//Explanantion:
	//line 1: a slice arr is created. value = [1 2 3 4]
	//line 2: s1 refers to arr. value = [1 2]
	//line 3: s2 also refers to arr. value = [2 3]
	//line 4: we append 10 on s1. value = [1 2 10].
	//line 5: prints arr, s1 and s2.
	// Since s1 is refering to arr, we get the value of arr = [1 2 10 4] and value of s2 = [2 10].
	//arr is a slice which is an underlying array in the heap.
	//when we assign arr[:2] to s1 then s1 refers to the underlying array arr instead of creating a whole new slice.
	//same in case of s2.
}
func q2() {
	//2. Reslicing and Capacity Growth

	// Problem:

	// What will be the output of the following Go program?
	arr := []int{1, 2, 3, 4, 5}
	s := arr[1:3]
	fmt.Println(len(s), cap(s)) // Length and capacity before appending
	s = append(s, 10, 20, 30)
	fmt.Println(arr, s)
	//Output:
	//2 4
	//[1 2 3 4 5] [2 3 10 20 30]
	//Explanation:
	//line 1: a slice arr is created. value = [1 2 3 4 5]
	//line 2: s refers to arr. value = [2 3]
	//line 3: prints the length of s and the capacity of s. len = 2  cap = 4
	//line 4: we append 10,20 and 30 in s. value = [2 3 10 20 30]
	//line 5: prints arr and s
	//here, the values of arr are not changed because we are appending
	//more values in the s (which is refering to arr) than arr was initially storing
	//therefore a new underlying array was created for s in the heap.
}

func modifySlice(s []int) {
	s[0] = 100
}

func q3() {
	// 3. Slice Reference Issue

	// Problem:

	// Why does the following Go program produce an unexpected result?
	arr := []int{1, 2, 3, 4, 5}
	s := arr[:3]
	modifySlice(s)
	fmt.Println(arr)
	//Output: [100 2 3 4 5]
	//Explanation:
	//line 1: a slice arr is created. value = [1 2 3 4 5]
	//line 2: s refers to arr. value = [2 3]
	//line 3: function modifySlice: s[0] = 100
	//here, since s is refering to slice arr
	//any changes made on s will be shown on arr
	//line4: prints the slice arr. value= [100 2 3 4 5]
}

func q4() {
	// 	4. Copying vs. Appending

	// Problem:

	// What will be printed when you run the following Go program?
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	s1[0] = 100
	fmt.Println(s1, s2)
	s1 = append(s1, 200)
	fmt.Println(s1, s2)
	//Output:
	//[100 2 3 4] [1 2 3 4]
	//[100 2 3 4 200] [1 2 3 4]
	//Explanation:
	//line1: creating a slice s1 with values [1 2 3 4]
	//line2: creating another slice s2 with length and capacity len(s1) i.e, 4
	//line3: copying all the data from s1 to s2.
	//line4: changing the element in index 0 of s1 from 1 to 100.
	//line5: printing s1 and s2. [100 2 3 4] [1 2 3 4]
	//line6: appending 200 on s1.
	//line7: printing s1 and s2 again. [100 2 3 4 200] [1 2 3 4]
	//here, any changes made in s1 was not shown in s2 because the data was simply copied
	//i.e, the s1 was not being referred by s2. Therefore, the changes were made only on s1.
}

func q5() {
	// 	5. Slice Capacity Overflow

	// Problem:

	// Predict the output and explain why the slice behaves this way.
	s := make([]int, 2, 3)
	s[0], s[1] = 10, 20
	s1 := append(s, 30)
	s2 := append(s, 40)
	fmt.Println(s, s1, s2)
	//Output:
	//[10 20] [10 20 40] [10 20 40]
	//Explanation:
	//line1: creating a slice s with length 2 and capacity 3
	//line2: storing values 10 and 20 in the index 0 and 1 of s respectively. [10 20]
	//line3: appending 30 on s and storing the value on s1. [10 20 30]
	//line4: appending 40 on s and storing the value on s2. [10 20 40]
	//line5: printing s, s1 and s2. [10 20] [10 20 40] [10 20 40]
	//here, we created an underlying array s of length 2 and capacity 3.
	//s1 on the other hand is referring to s, and when we append 30 on s and storing it on s1
	//the referred underlying array will get changed from [10 20] to [10 20 30]
	//For s the length is 2, therefore it will only see 2 elements in the underlying array.
	//As for s1 the length is 3, therefore it will see 3 elements in the underlying array.
	//And when we append 40 on s the underlying array will be changed from [10 20 30] to [10 20 40]
	//because 40 is being appended to s [10 20] and s2 is also referencing to the same underlying array.
	//for s2 also, the length is 3, therefore it will see 3 elements in the underlying array.
	//Since the length of s, s1 and s2 are 2, 3 and 3 respectively.
	//Therefore when we try to print the underlying array [10 20 40] we will get
	//s = [10 20], s1 = [10 20 40], s2 = [10 20 40]
}

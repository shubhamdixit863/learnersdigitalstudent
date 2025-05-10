package main

import "fmt"

func main() {

	// Write the Output and Explain Why?
	problem1()
	problem2()
	problem3()
	problem4()
	problem5()
}

func problem1() {
	arr := []int{1, 2, 3, 4} // Array of length 4 and Capacity 4 is created in heap with its slice descriptor in stack

	s1 := arr[:2] // Another slice descriptor S1 is created in stack pointing to the 0th index of same underlying array arr. (l=2, cap=4)

	s2 := arr[2:3] // Another slice descriptor S2 is created in stack pointing to the 2nd index of same underlying array arr. (l=1, cap=2)

	fmt.Println(s2) // S2 = [3]

	fmt.Println(&arr[0], &s1[0], &s2[0]) // Address of arr and S1 is same and s2 points to 2nd index of arr.

	s1 = append(s1, 10) // Since S1 has the same underlying array as arr and S2, when 10 is appended in S1 in 3rd pos, l will become 3 and arr and s2 will also have 3 there.

	arr = append(arr, 5) // Since arr has l=4 and cap=4, to append 5, a new underlying arr is created with new memory add of l=5 and cap=8. that add is passed to the pointer of arr descriptor.
	s2[0] = 99           // Now arr has new array but S1 and S2 still have the same underlying array, both S1 and S2 will change..

	fmt.Println(&arr[0], &s1[0], &s2[0])
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	fmt.Println(len(arr), cap(arr))
	fmt.Println(arr, s1, s2)
}

func problem2() {

	arr := []int{1, 2, 3, 4, 5} // Array of length 5 and Capacity 5 is created in heap with its slice descriptor in stack

	s := arr[1:3] //  Creates a slice s that refers to [2, 3] from arr(same underlying array).

	fmt.Println(len(s), cap(s)) // Length and capacity before appending will be l=2 and cap=4, Since it starts with 1st index of arr, cap will be 1 less then arr.
	// s has length = 2 (elements [2,3]) and capacity = 4 (from index 1 to end of arr).

	s = append(s, 10, 20, 30) // Since s still shares memory with arr, appending up to 2 elements (10, 20) modifies arr (because it fits within the capacity).
	// But appending 30 exceeds the slice’s capacity, causing a new underlying array to be allocated.
	// As a result, arr may be modified partially, but s points to a new array.

	fmt.Println(arr, s)

}

func modifySlice(s []int) {

	s[0] = 100 // This modifies s[0], but since s and arr refers to same underlying array, arr[0] is also changed.

}

func problem3() {

	arr := []int{1, 2, 3, 4, 5} // Array of length 5 and Capacity 5 is created in heap with its slice descriptor in stack

	s := arr[:3] // Creates a slice referring to [1, 2, 3] (first three elements of arr).
	// Slices are references to the same memory as arr.

	modifySlice(s) // Modifying a slice modifies the original array if it shares the same memory.

	fmt.Println(arr) // [100 2 3 4 5]

}

func problem4() {

	s1 := []int{1, 2, 3, 4} // Array of length 4 and Capacity 4 is created in heap with its slice descriptor in stack

	s2 := make([]int, len(s1)) // Another Array of length 4 and Capacity 4 is Created in Heap with its slice descriptor in stack, All the elementa are Zero.

	copy(s2, s1) // copies the elements from s1 into s2.
	// This ensures that s1 and s2 do not share the same memory.

	s1[0] = 100 // s1[0] is changed to 100, but s2[0] remains 1 because they have separate underlyinng Array.

	fmt.Println(s1, s2)

	s1 = append(s1, 200) // s1 is now extended with 200, but s2 remains unchanged.
	// Since S1 has l=4 and cap=4, to append 200, a new underlying array is created with new memory add of l=5 and cap=8. that add is passed to the pointer of S1 descriptor.

	fmt.Println(s1, s2) // [100 2 3 4 200] [1 2 3 4]

}

func problem5() {

	s := make([]int, 2, 3) // creates a slice with length = 2, capacity = 3 with Zero as its elements.

	s[0], s[1] = 10, 20 // 10 and 20 is assigned to 0th and 1st index of the Underlying array.

	s1 := append(s, 30)           // A new slice descriptor is created which points to the same underlying array Appends 30, which fits within the capacity (3).
	fmt.Println(len(s1), cap(s1)) // l=3, cap=3
	s2 := append(s, 40)           // again A new slice descriptor is created which points to the same underlying array Appends 40 instead, which fits within the capacity (3). As length of s is 2
	// Since s and s1 refer to the same memory, s2 also initially shares it.
	// if we further append s, S1 or S2..it will allocate the seprate memory for that slice.
	fmt.Println(len(s2), cap(s2))      // l=3, cap=3
	fmt.Println(&s[0], &s1[0], &s2[0]) // 0xc0000160f0 0xc0000160f0 0xc0000160f0

	fmt.Println(s, s1, s2) // [10 20] [10 20 40] [10 20 40]

}

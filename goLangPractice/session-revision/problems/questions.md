1. Slice Capacity & Append Behavior

Problem:

What will be the output of the following Go program? Explain why.

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4}

	s1 := arr[:2]

	s2 := arr[1:3]

	s1 = append(s1, 10) // Modify slice s1

	fmt.Println(arr, s1, s2)

}

2. Reslicing and Capacity Growth



Problem:

What will be the output of the following Go program?

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}

	s := arr[1:3]

	fmt.Println(len(s), cap(s)) // Length and capacity before appending

	s = append(s, 10, 20, 30)

	fmt.Println(arr, s)

}

3. Slice Reference Issue



Problem:

Why does the following Go program produce an unexpected result?

package main

import "fmt"

func modifySlice(s []int) {

    s[0] = 100

}

func main() {

    arr := []int{1, 2, 3, 4, 5}

    s := arr[:3]

    modifySlice(s)

    fmt.Println(arr)
}

4. Copying vs. Appending



Problem:

What will be printed when you run the following Go program?

package main

import "fmt"

func main() {

    s1 := []int{1, 2, 3, 4}

    s2 := make([]int, len(s1))

    copy(s2, s1)

    s1[0] = 100

    fmt.Println(s1, s2)
    s1 = append(s1, 200)

    fmt.Println(s1, s2)

}

5. Slice Capacity Overflow



Problem:

Predict the output and explain why the slice behaves this way.

package main

import "fmt"

func main() {

    s := make([]int, 2, 3)

    s[0], s[1] = 10, 20

    s1 := append(s, 30)

    s2 := append(s, 40)

    fmt.Println(s, s1, s2)

}

Bonus Challenge



Create a function that mimics append() by manually reallocating memory when needed.

func myAppend(s []int, elems ...int) []int {

    // Implement capacity expansion logic similar to Go's built-in append()

}


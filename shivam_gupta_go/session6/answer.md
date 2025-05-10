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



answer : [1,2,10,4]   [1,2,10]  [2,10]
Appending to a slice in Go modifies the underlying array if capacity allows, affecting other slices referencing it. Otherwise, a new array is allocated, leaving the original unchanged.

****************************************************************************************************************************


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


answer : output [1,2,3,4,5] [2,3,10,20,30]
Appending beyond slice capacity creates a new array, leaving the original unchanged.

*******************************************************************************************************

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
answer : output [100,2,3]
slice -> pass by  refrence 

***********************************************************************************************

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
answer : output: [100,2,3,4] [1,2,3,4],  [100,2,3,4,200] [1,2,3,4]
Copy creates a separate slice; changes to s1 don't affect s2. Append extends s1, leaving s2 unchanged

******************************************************************************************************************


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
answer : output : [10,20] [10,20,40] [10,20,40]
Appending beyond capacity creates a new underlying array, causing s1 and s2 to share the same data.
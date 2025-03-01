Que:1 
    package main

    import "fmt"

    func main() {

        arr := []int{1, 2, 3, 4}

        s1 := arr[:2]

        s2 := arr[1:3]

        s1 = append(s1, 10) // Modify slice s1

        fmt.Println(arr, s1, s2)

    }

ANS: [1 2 10 4] [1 2 10] [2 10]
// when the slice is changed till its capacity then the underlying array also changes.

QUE:2
    package main

    import "fmt"

    func main() {

        arr := []int{1, 2, 3, 4, 5}

        s := arr[1:3]

        fmt.Println(len(s), cap(s)) // Length and capacity before appending

        s = append(s, 10, 20, 30)

        fmt.Println(arr, s)

    }

ANS: 2 4
    [1 2 3 4 5] [2 3 10 20 30]
// when the slice exceeds its capacity the go create new slice with double capacity ad in that case underlying array do not change.

Que:3 
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

ANS: [100 2 3 4 5]
// Because slice changes so the underlying array also changes.

Que:4 
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

ANS: [100 2 3 4] [1 2 3 4]
    [100 2 3 4 200] [1 2 3 4]

Que:5
    package main

    import "fmt"

    func main() {

        s := make([]int, 2, 3)

        s[0], s[1] = 10, 20

        s1 := append(s, 30)

        s2 := append(s, 40)

        fmt.Println(s, s1, s2)

    }

ANS: [10 20] [10 20 40] [10 20 40]

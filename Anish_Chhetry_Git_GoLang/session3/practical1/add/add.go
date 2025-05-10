package add

import "fmt"

func AddStudent(names []string, scores []int) ([]string, []int) {
	var name string
	var marks int
	fmt.Println("Enter Student Name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Enter Score: ")
	fmt.Scanf("%d\n", &marks)
	names = append(names, name)
	scores = append(scores, marks)
	fmt.Println("Student added successfully!")

	return names, scores
}

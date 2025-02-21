package update

import "fmt"

func UpdateStudent(names []string, scores []int) ([]string, []int) {
	var name string
	var marks int
	fmt.Println("Enter Student Name to Update: ")
	fmt.Scanf("%s\n", &name)
	for i := 0; i <= len(names); i++ {
		if names[i] == name {
			fmt.Println("Enter New Score: ")
			fmt.Scanf("%d\n", &marks)
			scores[i] = marks
			fmt.Println("Score updated successfully!")
			return names, scores
		}
		fmt.Println("Student not found")
	}

	return names, scores
}

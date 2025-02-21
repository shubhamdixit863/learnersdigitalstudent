package delete

import "fmt"

func DeleteStudent(names []string, scores []int) ([]string, []int) {
	var name string
	fmt.Println("Enter Student Name to Delete: ")
	fmt.Scanf("%s\n", &name)
	for i := 0; i <= len(names); i++ {
		if names[i] == name {
			names = append(names[:i], names[i+1:]...)
			scores = append(scores[:i], scores[i+1:]...)
			fmt.Println("Student deleted successfully!")
			return names, scores
		}
		fmt.Println("Student not found")
	}
	return names, scores
}

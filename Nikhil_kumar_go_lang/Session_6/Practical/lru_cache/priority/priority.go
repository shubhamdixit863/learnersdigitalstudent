package priority

import "fmt"

func Change_priority(set_priority int, priority []int) []int {
	for i := 0; i < len(priority); i++ {
		if priority[i] == set_priority {
			priority = append(priority[:i], priority[i+1:]...)
			break
		}
	}
	priority = append(priority, set_priority)
	return priority
}

func Remove_priority(priority []int) []int {
	if len(priority) > 0 {
		priority = priority[1:]
		return priority
	}
	return priority
}

func Show_priority(priority []int) {
	fmt.Println(priority)
}

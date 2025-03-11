package services

import "fmt"

type Taskmanager struct {
	TaskDone []string
}

func (t *Taskmanager) NewTask() {

	fmt.Println("1. File Processing\n 2. Apicall\n 3. Datavalidation\n 4. View Task Done")
	var choose int
	fmt.Scanln(&choose)
	switch choose {
	case 1:
		o1 := NewFile()
		fmt.Println(o1.Run())
		t.TaskDone = append(t.TaskDone, "fileprocessing")
	case 2:
		o1 := NewApiCall()
		fmt.Println(o1.Run())
		t.TaskDone = append(t.TaskDone, "apicall")
	case 3:
		o1 := NewDataValidation()
		fmt.Println(o1.Run())
		t.TaskDone = append(t.TaskDone, "datavalidation")
	case 4:
		fmt.Println(t.TaskDone)

	}

}

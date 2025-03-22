package services

import (
	"fmt"
	"ResultManagement/internal/models"
)


func ProcessResults(students []models.Grade) {
	for _, student := range students {
		fmt.Println(student.GetDetails())
	}
}

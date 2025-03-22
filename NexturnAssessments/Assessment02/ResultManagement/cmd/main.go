package main

import (
	"ResultManagement/internal/models"
	"ResultManagement/internal/services"
)

func main() {
	
	engStudent := models.EngineeringStudent{
		Student: models.Student{Name: "Sanskruti", ID: 101, Stream: "Engineering"},
		Electronics:95,
		Signals: 85,
		Vlsi: 98,
	}

	artsStudent := models.ArtsStudent{
		Student: models.Student{Name: "Kalyani", ID: 102, Stream: "Arts"},
		History: 90,
		Geography: 90,
		Civics: 85,
	}

	
	students := []models.Grade{&engStudent, &artsStudent}
	services.ProcessResults(students)
}

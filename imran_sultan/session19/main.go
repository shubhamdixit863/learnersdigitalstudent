package main

import (
	"html/template"
	"net/http"
)

// Person Struct
type Person struct {
	ID          int
	Name        string
	PhoneNumber string
}

var people []Person // Slice to store people
var idCounter = 1   // Auto-increment ID

// Home Page Handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/static/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Create Page Handler
func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		phone := r.FormValue("phone")

		newPerson := Person{
			ID:          idCounter,
			Name:        name,
			PhoneNumber: phone,
		}
		idCounter++
		people = append(people, newPerson)

		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("internal/static/create.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// List Page Handler
func listHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/static/list.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, people)
	// file, err := os.ReadFile("internal/static/list.html")
	// if err != nil {
	// 	return

	// }
	// w.Write(file)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("internal/static"))) // Home page
	http.HandleFunc("/create", createHandler)                      // Create page
	http.HandleFunc("/list", listHandler)                          // List page
	// http.Handle("/list", http.FileServer(http.Dir("internal/static/list.html"))) // List page

	http.ListenAndServe(":8085", nil)
}

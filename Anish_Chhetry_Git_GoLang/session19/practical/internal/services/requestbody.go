package services

type RequestBody struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

var users = []RequestBody{}

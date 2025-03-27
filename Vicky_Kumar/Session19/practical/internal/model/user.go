package model

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(id string, name string, email string) *User {
	return &User{Id: id, Name: name, Email: email}
}

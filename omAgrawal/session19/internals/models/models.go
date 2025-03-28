package models

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   string `json:"id"`
}

func NewUser() *User {
	return &User{}

}

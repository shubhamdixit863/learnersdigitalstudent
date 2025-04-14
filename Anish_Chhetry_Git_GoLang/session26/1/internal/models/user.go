package models

// model which will be stored in Db
type User struct {
	Username   string
	Password   string
	FirstName  string
	SecondName string
	Email      string
	ID         string
}

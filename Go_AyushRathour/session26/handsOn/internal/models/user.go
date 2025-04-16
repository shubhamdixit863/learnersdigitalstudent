package models

//type User struct {
//	Username   string
//	Password   string
//	FirstName  string
//	SecondName string
//	Email      string
//	ID         string
//}

type User struct {
	Username   string `gorm:"column:UserName"`
	Password   string `gorm:"column:password"`
	FirstName  string `gorm:"column:FirstName"`
	SecondName string `gorm:"column:SecondName"`
	ID         string `gorm:"primaryKey"`
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "Users"
}

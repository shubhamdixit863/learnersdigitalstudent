package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"session26/internal/models"
)

func Migrate() {

	dsn := "root:root@tcp(127.0.0.1:3306)/myapp?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	err = db.AutoMigrate(&models.Products{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Database migrated !!!")
}

func main() {
	Migrate()
}

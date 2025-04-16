package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"session26/internal/models"
)

type MysqlOrm struct {
	db *gorm.DB
}

func (m MysqlOrm) CreateUser(ctx context.Context, user models.User) (interface{}, error) {
	tx := m.db.Create(&user)
	if tx.Error != nil {
		log.Println("error", tx.Error)
		return "", tx.Error
	}

	return user.ID, nil
}

func (m MysqlOrm) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	var user models.User
	log.Println("Incoming username", userName)
	tx := m.db.Where("UserName = ?", userName).Find(&user)
	if tx.Error != nil {
		log.Println("error", tx.Error)
		return nil, tx.Error
	}
	if user.Username == "" {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

func (m MysqlOrm) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := m.db.WithContext(ctx).Find(&users).Error; err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	return users, nil
}

func (m MysqlOrm) UpdateUser(ctx context.Context, user models.User) error {
	if err := m.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		log.Println("Error updating user:", err)
		return err
	}
	return nil
}

func (m MysqlOrm) DeleteUser(ctx context.Context, user models.User) error {
	if err := m.db.WithContext(ctx).Where("id = ?", user.ID).Delete(&models.User{}).Error; err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	return nil
}

func NewMysqlOrm(db *gorm.DB) DbRepository {
	return &MysqlOrm{db: db}
}

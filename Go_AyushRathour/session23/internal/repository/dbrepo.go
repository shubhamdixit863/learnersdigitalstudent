package repository

import "JWT/internal/models"

type DbRepository interface {
	CreateUser(user models.User) string
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}

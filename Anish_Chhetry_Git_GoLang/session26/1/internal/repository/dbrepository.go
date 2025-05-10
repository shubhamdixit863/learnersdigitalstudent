package repository

import (
	"1/internal/models"
	"context"
)

type DbRepository interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
	GetUserByUserName(ctx context.Context, userName string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, user models.User) error
}

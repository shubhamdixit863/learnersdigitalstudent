package repository

import (
	"assesment_5/internal/models"
)

type DbRepository interface {
	GetAll() []models.Post
	Create(models.Post) models.Post
	Update(int, models.Post) (models.Post, error)
	Delete(int) error
}

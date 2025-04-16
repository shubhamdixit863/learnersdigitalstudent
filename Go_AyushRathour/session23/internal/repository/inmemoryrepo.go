package repository

import (
	"JWT/internal/models"
	"errors"
)

type InMemoryRepository struct {
	users map[string]models.User
}

func NewInMemory() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[string]models.User),
	}
}

func (i *InMemoryRepository) GetUserByUsername(username string) (*models.User, error) {
	if v, ok := i.users[username]; ok {
		return &v, nil
	}
	return nil, errors.New("user  Not Found")
}

func (i *InMemoryRepository) GetAllUsers() ([]*models.User, error) {
	var slc []*models.User
	for _, v := range i.users {
		slc = append(slc, &v)
	}
	return slc, nil
}

// creating a new user here
func (i *InMemoryRepository) CreateUser(user models.User) string {
	i.users[user.Username] = user
	return user.ID
}

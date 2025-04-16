package repository

import (
	"context"
	"errors"
	"session26/internal/models"
	"sync"
)

type InMemoryRepository struct {
	users map[string]models.User
	mu    sync.RWMutex
}

func (i *InMemoryRepository) CreateUser(ctx context.Context, user models.User) (interface{}, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if _, exists := i.users[user.Username]; exists {
		return nil, errors.New("user already exists")
	}
	i.users[user.Username] = user
	return user.ID, nil
}

func (i *InMemoryRepository) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	user, exists := i.users[userName]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (i *InMemoryRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	var users []*models.User
	for _, user := range i.users {
		u := user
		users = append(users, &u)
	}
	return users, nil
}

func (i *InMemoryRepository) UpdateUser(ctx context.Context, user models.User) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if _, exists := i.users[user.Username]; !exists {
		return errors.New("user not found")
	}
	i.users[user.Username] = user
	return nil
}

func (i *InMemoryRepository) DeleteUser(ctx context.Context, user models.User) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if _, exists := i.users[user.Username]; !exists {
		return errors.New("user not found")
	}
	delete(i.users, user.Username)
	return nil
}

func NewInMemory() DbRepository {
	return &InMemoryRepository{
		users: make(map[string]models.User),
	}
}

//func (i *InMemoryRepository) GetUserByUsername(username string) (*models.User, error) {
//	if v, ok := i.users[username]; ok {
//		return &v, nil
//	}
//	return nil, errors.New("user  Not Found")
//}

//func (i *InMemoryRepository) GetAllUsers() ([]*models.User, error) {
//	var slc []*models.User
//	for _, v := range i.users {
//		slc = append(slc, &v)
//	}
//	return slc, nil
//}

//// creating a new user here
//func (i *InMemoryRepository) CreateUser(user models.User) interface{} {
//	i.users[user.Username] = user
//	return user.ID
//}

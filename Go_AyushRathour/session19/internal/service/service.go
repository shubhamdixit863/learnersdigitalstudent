package service

import (
	"crud/internal/model"
	"crud/internal/storage"
	"errors"
)

// UserService handles user operations
type UserService struct {
	store *storage.UserStore
}

// NewUserService initializes a new UserService
func NewUserService(store *storage.UserStore) *UserService {
	return &UserService{store: store}
}

// RegisterUser adds a new user
func (s *UserService) RegisterUser(user model.User) (model.User, error) {
	user.ID = s.store.NextID
	s.store.Users[user.ID] = user
	s.store.NextID++
	return user, nil
}

// GetUser fetches a user by ID
func (s *UserService) GetUser(id int) (model.User, error) {
	user, exists := s.store.Users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

// GetAllUsers fetches all users
func (s *UserService) GetAllUsers() map[int]model.User {
	return s.store.Users
}

// UpdateUser updates user details
func (s *UserService) UpdateUser(id int, updatedUser model.User) (model.User, error) {
	_, exists := s.store.Users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	updatedUser.ID = id
	s.store.Users[id] = updatedUser
	return updatedUser, nil
}

// DeleteUser removes a user
func (s *UserService) DeleteUser(id int) error {
	_, exists := s.store.Users[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(s.store.Users, id)
	return nil
}

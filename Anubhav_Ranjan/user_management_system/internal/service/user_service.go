package service

import (
	"user_management_system/internal/model"
	"user_management_system/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (service *UserService) CreateUser(user *model.User) (*model.User, error) {
	// Business logic goes here
	return service.UserRepo.CreateUser(user)
}

func (service *UserService) GetUser(id int) (*model.User, error) {
	return service.UserRepo.GetUser(id)
}

func (service *UserService) GetAllUsers() ([]model.User, error) {
	return service.UserRepo.GetAllUsers()
}

func (service *UserService) UpdateUser(user *model.User) (*model.User, error) {
	return service.UserRepo.UpdateUser(user)
}

func (service *UserService) DeleteUser(id int) error {
	return service.UserRepo.DeleteUser(id)
}

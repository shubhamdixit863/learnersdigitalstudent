package services

import (
	"errors"
	"practical/internal/model"
	"practical/internal/repositories"
	"practical/internal/utils"
)

func AddUser(user model.User) error {
	repositories.Users = append(repositories.Users, user)
	return nil
}

func GetUsers() ([]model.User, error) {
	return repositories.Users, nil
}

func GetUserById(id string) (model.User, error) {
	for _, u := range repositories.Users {
		if u.Id == id {
			return u, nil
		}
	}
	return model.User{}, errors.New(utils.UserNotFoundErr)
}

func UpdateUser(id string, updatedUser model.User) (model.User, error) {
	for i, u := range repositories.Users {
		if u.Id == id {
			repositories.Users[i].Name = updatedUser.Name
			repositories.Users[i].Email = updatedUser.Email
			return repositories.Users[i], nil
		}
	}
	return model.User{}, errors.New(utils.UserNotFoundErr)
}

func DeleteUser(id string) error {
	for i, u := range repositories.Users {
		if u.Id == id {
			repositories.Users = append(repositories.Users[:i], repositories.Users[i+1:]...)
			return nil
		}
	}
	return errors.New(utils.UserNotFoundErr)
}

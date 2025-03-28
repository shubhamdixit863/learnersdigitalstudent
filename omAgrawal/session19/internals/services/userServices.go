package services

import (
	"encoding/json"
	"errors"
	"session19/internals/models"
)

type SliceOfUser struct {
	Slice []*models.User
}

func NewUserSlice() *SliceOfUser {
	return &SliceOfUser{}
}

func (Sl *SliceOfUser) RegisterUSer(new *models.User) {
	Sl.Slice = append(Sl.Slice, new)

}
func (Sl *SliceOfUser) GetAllUser() ([]byte, error) {
	data, err := json.Marshal(Sl)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}

}

func (Sl *SliceOfUser) GetUserById(id string) ([]byte, error) {
	for _, v := range Sl.Slice {
		if v.Id == id {
			data, err := json.Marshal(&v)
			if err != nil {
				return nil, err
			} else {
				return data, nil
			}
		}
	}
	return nil, errors.New("id not Found")
}

func (Sl *SliceOfUser) UpdateUserById(id string, new *models.User) error {
	var x int
	for i, v := range Sl.Slice {
		if v.Id == id {
			v = new
			x = i
			return nil
		}
	}
	Sl.Slice[x] = new
	return errors.New("id not Found")
}

func (Sl *SliceOfUser) DeleteUserById(id string) error {
	for i, v := range Sl.Slice {
		if v.Id == id {
			Sl.Slice = append(Sl.Slice[:i], Sl.Slice[i+1:]...)
		}
	}
	return errors.New("id not Found")
}

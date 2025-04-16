package storage

import "crud/internal/model"

// UserStore stores user data
type UserStore struct {
	Users  map[int]model.User
	NextID int
}

// NewUserStore initializes a new UserStore
func NewUserStore() *UserStore {
	return &UserStore{
		Users:  make(map[int]model.User),
		NextID: 1,
	}
}

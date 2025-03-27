package storage

import "crud/internal/model"

// UserStore stores user data in memory
type UserStore struct {
	Users  map[int]model.User
	NextID int
}

// NewUserStore initializes a new user store
func NewUserStore() *UserStore {
	return &UserStore{
		Users:  make(map[int]model.User),
		NextID: 1,
	}
}

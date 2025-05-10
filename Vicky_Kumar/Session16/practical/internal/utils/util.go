package utils

import "github.com/google/uuid"

func RandomID() string {
	id := uuid.New()
	return id.String()
}

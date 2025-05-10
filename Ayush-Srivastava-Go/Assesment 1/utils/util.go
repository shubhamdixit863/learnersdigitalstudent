package utils

import "math/rand"

func GenerateShortCode() string {

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortCode := ""

	for range 6 {
		shortCode += string(letters[rand.Intn(len(letters))])
	}

	return shortCode

}
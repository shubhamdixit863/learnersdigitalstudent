package handlers

import (
	"practical/internal/repository"
	"practical/services"
)

type Handler struct {
	repo       repository.DbRepository
	jwtService *services.JWTService
}

func NewHandler(repo repository.DbRepository, jwtService *services.JWTService) *Handler {
	return &Handler{
		repo:       repo,
		jwtService: jwtService,
	}
}

package handlers

import (
	"JWT/internal/repository"
	"JWT/internal/services"
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

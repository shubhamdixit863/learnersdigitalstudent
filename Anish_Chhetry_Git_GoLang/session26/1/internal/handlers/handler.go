package handlers

import (
	"1/internal/repository"
	"1/internal/services"
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

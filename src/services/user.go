package services

import (
	"context"
	"github.com/delos/aquafarm-management/src/repositories"

	"github.com/delos/aquafarm-management/src/models"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository}
}

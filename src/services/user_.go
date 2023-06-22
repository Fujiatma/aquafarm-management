package services

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
)

func (u userService) CreateUser(ctx context.Context, user *models.User) error {
	user.ID = models.GenerateID()
	err := u.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

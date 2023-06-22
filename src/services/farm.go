package services

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/delos/aquafarm-management/src/repositories"
)

type FarmService interface {
	CreateFarm(ctx context.Context, farm *models.Farm) error
	UpdateFarm(ctx context.Context, farm *models.Farm) (*models.Farm, error)
	GetFarmByID(ctx context.Context, farmID string) (farm *models.Farm, err error)
	GetAllFarmByUserID(ctx context.Context, userID string) (farms []*models.Farm, err error)
	GetFarmByName(ctx context.Context, farmName string) (*models.Farm, error)
	DeleteFarm(ctx context.Context, farm *models.Farm) error
}

type farmService struct {
	farmRepository repositories.FarmRepository
}

func NewFarmService(farmRepository repositories.FarmRepository) FarmService {
	return &farmService{farmRepository}
}

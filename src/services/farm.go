package services

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/delos/aquafarm-management/src/repositories"
)

type FarmService interface {
	CreateFarm(ctx context.Context, farm models.Farm) error
	UpdateFarm(ctx context.Context, farm models.Farm) error
	GetFarmByID(ctx context.Context, farmID string) (farm *models.Farm, err error)
	GetAllFarms(ctx context.Context) (farms []models.Farm, err error)
}

type farmService struct {
	farmRepository repositories.FarmRepository
}

func NewFarmsService(farmRepository repositories.FarmRepository) FarmService {
	return &farmService{farmRepository}
}

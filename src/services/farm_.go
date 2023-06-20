package services

import (
	"context"

	"github.com/delos/aquafarm-management/src/models"
)

func (s farmService) CreateFarm(ctx context.Context, farm models.Farm) error {
	return s.farmRepository.CreateFarm(ctx, farm)
}

func (s farmService) UpdateFarm(ctx context.Context, farm models.Farm) error {
	//TODO implement me
	panic("implement me")
}

func (s farmService) GetFarmByID(ctx context.Context, farmID string) (farm *models.Farm, err error) {
	//TODO implement me
	panic("implement me")
}

func (s farmService) GetAllFarms(ctx context.Context) (farms []models.Farm, err error) {
	//TODO implement me
	panic("implement me")
}

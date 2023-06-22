package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/delos/aquafarm-management/src/models"
	"gorm.io/gorm"
	"time"
)

func (s farmService) CreateFarm(ctx context.Context, farm *models.Farm) error {
	// Check if farm with the same name already exists
	existingFarm, err := s.farmRepository.GetFarmByName(ctx, farm.FarmName)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if existingFarm != nil {
		return fmt.Errorf("farm with the same name already exists")
	}

	// Create the farm
	farm.ID = models.GenerateID()
	err = s.farmRepository.CreateFarm(ctx, farm)
	if err != nil {
		return err
	}

	return nil
}

func (s farmService) UpdateFarm(ctx context.Context, farmData *models.Farm) (*models.Farm, error) {
	farm, err := s.farmRepository.GetFarmByID(ctx, farmData.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jika farm tidak ditemukan, buat farm baru dengan ID yang diberikan
			farmData.ID = models.GenerateID()
			farmData.CreatedAt = time.Now()
			err := s.CreateFarm(ctx, farmData)
			if err != nil {
				return nil, err
			}
			return farmData, nil
		}
		return nil, err
	}

	// Jika farm ditemukan, update farm yang ada dengan data baru
	farm.UpdatedAt = farmData.UpdatedAt
	farm.FarmName = farmData.FarmName
	err = s.farmRepository.UpdateFarm(ctx, farm)
	if err != nil {
		return nil, err
	}

	return farmData, nil
}

func (s farmService) DeleteFarm(ctx context.Context, payload *models.Farm) error {
	farm, err := s.farmRepository.GetFarmByID(ctx, payload.ID)
	if err != nil {
		return err
	}

	if farm.UserID != payload.UserID {
		return fmt.Errorf("user does not have permission to delete the farm")
	}

	return s.farmRepository.SoftDeleteFarm(ctx, farm)
}

func (s farmService) GetFarmByID(ctx context.Context, farmID string) (farm *models.Farm, err error) {
	farm, err = s.farmRepository.GetFarmByID(ctx, farmID)
	if err != nil {
		return nil, err
	}

	return farm, nil
}

func (s farmService) GetAllFarmByUserID(ctx context.Context, userID string) (farms []*models.Farm, err error) {
	farms, err = s.farmRepository.GetFarmsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return farms, nil
}

func (s farmService) GetFarmByName(ctx context.Context, farmName string) (*models.Farm, error) {
	farm, err := s.farmRepository.GetFarmByName(ctx, farmName)
	if err != nil {
		return nil, err
	}
	return farm, nil
}

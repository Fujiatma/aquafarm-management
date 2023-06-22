package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"

	"github.com/delos/aquafarm-management/src/models"
)

func (s pondService) CreatePond(ctx context.Context, pond *models.Pond) error {
	// Check if pond with the same name already exists
	existingPond, err := s.pondRepository.GetPondByName(ctx, pond.PondName)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if existingPond != nil {
		return fmt.Errorf("pond with the same name already exists")
	}

	// Check if farm with the specified ID exists
	farm, err := s.farmRepository.GetFarmByID(ctx, pond.FarmID)
	if err != nil {
		return err
	}
	if farm == nil {
		return fmt.Errorf("specified farm does not exist")
	}

	// Create the pond
	pond.ID = models.GenerateID()
	err = s.pondRepository.CreatePond(ctx, pond)
	if err != nil {
		return err
	}

	pond.Farm = *farm

	return nil
}

func (s pondService) UpdatePond(ctx context.Context, pondData *models.Pond) (*models.Pond, error) {
	existingPond, err := s.pondRepository.GetPondByID(ctx, pondData.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jika pond tidak ditemukan, buat pond baru dengan ID yang diberikan
			pondData.CreatedAt = time.Now()
			pondData.ID = models.GenerateID()
			err := s.CreatePond(ctx, pondData)
			if err != nil {
				return nil, err
			}
			return pondData, nil
		}
		return nil, err
	}

	// Jika pond ditemukan, update pond yang ada dengan data baru
	existingPond.PondName = pondData.PondName
	err = s.pondRepository.UpdatePond(ctx, existingPond)
	if err != nil {
		return nil, err
	}

	return pondData, nil
}

func (s pondService) DeletePond(ctx context.Context, payload *models.Pond, userID string) error {
	pond, err := s.pondRepository.GetPondByID(ctx, payload.ID)
	if err != nil {
		return err
	}

	if pond.Farm.UserID != userID {
		return fmt.Errorf("user does not have permission to delete the pond")
	}

	return s.pondRepository.SoftDeletePond(ctx, pond)
}

func (s pondService) GetPondByID(ctx context.Context, pondID string) (pond *models.Pond, err error) {
	pond, err = s.pondRepository.GetPondByID(ctx, pondID)
	if err != nil {
		return nil, err
	}

	return pond, nil
}

func (s pondService) GetPondsByUserID(ctx context.Context, userID string) (ponds []*models.Pond, err error) {
	ponds, err = s.pondRepository.GetPondsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return ponds, nil
}

func (s pondService) GetPondByName(ctx context.Context, pondName string) (pond *models.Pond, err error) {
	pond, err = s.pondRepository.GetPondByName(ctx, pondName)
	if err != nil {
		return nil, err
	}
	return pond, nil
}

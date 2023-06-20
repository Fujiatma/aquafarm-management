package repositories

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
	"gorm.io/gorm"
)

type FarmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) *FarmRepository {
	return &FarmRepository{db: db}
}

func (r *FarmRepository) CreateFarm(ctx context.Context, farm models.Farm) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(farm).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (r *FarmRepository) UpdateFarm(ctx context.Context, farm models.Farm) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(farm).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *FarmRepository) GetFarmByID(ctx context.Context, farmID string) (farm *models.Farm, err error) {
	err = r.db.WithContext(ctx).First(farm, "id = ?", farmID).Error
	if err != nil {
		return nil, err
	}
	return farm, nil
}

func (r *FarmRepository) GetAllFarms(ctx context.Context) (farms []models.Farm, err error) {
	err = r.db.WithContext(ctx).Find(&farms).Error
	if err != nil {
		return nil, err
	}

	return farms, nil
}

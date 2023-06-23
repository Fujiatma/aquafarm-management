package repositories

import (
	"context"
	"errors"
	"github.com/delos/aquafarm-management/src/models"
	"gorm.io/gorm"
	"time"
)

type FarmRepository interface {
	CreateFarm(ctx context.Context, farm *models.Farm) error
	UpdateFarm(ctx context.Context, farm *models.Farm) error
	GetFarmByID(ctx context.Context, farmID string) (*models.Farm, error)
	GetFarmsByUserID(ctx context.Context, userID string) ([]*models.Farm, error)
	GetFarmByName(ctx context.Context, farmName string) (*models.Farm, error)
	SoftDeleteFarm(ctx context.Context, farm *models.Farm) error
}

type farmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) FarmRepository {
	return &farmRepository{db: db}
}

func (r *farmRepository) CreateFarm(ctx context.Context, farm *models.Farm) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(farm).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (r *farmRepository) UpdateFarm(ctx context.Context, farm *models.Farm) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&farm).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *farmRepository) GetFarmByID(ctx context.Context, farmID string) (farm *models.Farm, err error) {
	err = r.db.WithContext(ctx).First(&farm, "id = ?", farmID).Error
	if err != nil {
		return nil, err
	}
	return farm, nil
}

func (r *farmRepository) GetFarmsByUserID(ctx context.Context, userID string) (farms []*models.Farm, err error) {
	err = r.db.WithContext(ctx).Preload("Ponds").Where("user_id = ?", userID).Find(&farms).Error
	if err != nil {
		return nil, err
	}

	return farms, nil
}

func (r *farmRepository) GetFarmByName(ctx context.Context, farmName string) (farm *models.Farm, err error) {
	err = r.db.WithContext(ctx).Where("farm_name = ?", farmName).First(&farm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return farm, nil
}

func (r *farmRepository) SoftDeleteFarm(ctx context.Context, farm *models.Farm) error {
	now := time.Now()
	farm.IsDeleted = true
	farm.DeletedAt = &now

	return r.db.WithContext(ctx).Save(farm).Error
}

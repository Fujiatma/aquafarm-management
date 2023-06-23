package repositories

import (
	"context"
	"errors"
	"github.com/delos/aquafarm-management/src/models"
	"gorm.io/gorm"
	"time"
)

type PondRepository interface {
	CreatePond(ctx context.Context, pond *models.Pond) error
	UpdatePond(ctx context.Context, pond *models.Pond) error
	GetPondByID(ctx context.Context, pondID string) (*models.Pond, error)
	GetPondsByUserID(ctx context.Context, userID string) ([]*models.Pond, error)
	GetPondByName(ctx context.Context, pondName string) (*models.Pond, error)
	SoftDeletePond(ctx context.Context, pond *models.Pond) error
}

type pondRepository struct {
	db *gorm.DB
}

func NewPondRepository(db *gorm.DB) PondRepository {
	return &pondRepository{db: db}
}

func (r *pondRepository) CreatePond(ctx context.Context, pond *models.Pond) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(pond).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *pondRepository) UpdatePond(ctx context.Context, pond *models.Pond) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&pond).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *pondRepository) GetPondByID(ctx context.Context, pondID string) (pond *models.Pond, err error) {
	err = r.db.WithContext(ctx).Preload("Farm").First(&pond, "id = ?", pondID).Error
	if err != nil {
		return nil, err
	}

	return pond, nil
}

func (r *pondRepository) GetPondsByUserID(ctx context.Context, userID string) (ponds []*models.Pond, err error) {
	err = r.db.WithContext(ctx).Joins("JOIN farms ON ponds.farm_id = farms.id").Where("farms.user_id = ?", userID).Find(&ponds).Error
	if err != nil {
		return nil, err
	}

	return ponds, nil
}

func (r *pondRepository) GetPondByName(ctx context.Context, pondName string) (pond *models.Pond, err error) {
	err = r.db.WithContext(ctx).Where("pond_name = ?", pondName).First(&pond).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return pond, nil
}

func (r *pondRepository) SoftDeletePond(ctx context.Context, pond *models.Pond) error {
	now := time.Now()
	pond.IsDeleted = true
	pond.DeletedAt = &now

	return r.db.WithContext(ctx).Save(pond).Error
}

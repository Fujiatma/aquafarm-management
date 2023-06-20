package repositories

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
	"gorm.io/gorm"
)

type PondRepository struct {
	db *gorm.DB
}

func NewPondRepository(db *gorm.DB) *PondRepository {
	return &PondRepository{db: db}
}

func (r *PondRepository) CreatePond(ctx context.Context, pond models.Pond) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(pond).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *PondRepository) UpdatePond(ctx context.Context, pond models.Pond) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(pond).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (r *PondRepository) GetPondByID(ctx context.Context, pondID string) (pond *models.Pond, err error) {
	err = r.db.WithContext(ctx).First(pond, "id = ?", pondID).Error
	if err != nil {
		return nil, err
	}

	return pond, nil
}

func (r *PondRepository) GetAllPonds(ctx context.Context) (ponds []*models.Pond, err error) {
	err = r.db.WithContext(ctx).Find(&ponds).Error
	if err != nil {
		return nil, err
	}

	return ponds, nil
}

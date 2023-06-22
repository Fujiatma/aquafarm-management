package repositories

import (
	"context"
	"errors"
	"github.com/delos/aquafarm-management/src/models"
	"gorm.io/gorm"
	"time"
)

type StatisticRepository interface {
	CreateStatistic(ctx context.Context, statistic models.Statistic) error
	GetAllStatistics(ctx context.Context) ([]*models.Statistic, error)
	UpsertStatistic(ctx context.Context, statistic *models.Statistic) error
}

type statisticRepository struct {
	db *gorm.DB
}

func NewStatisticRepository(db *gorm.DB) StatisticRepository {
	return &statisticRepository{db: db}
}

func (r *statisticRepository) CreateStatistic(ctx context.Context, statistic models.Statistic) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(statistic).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (r *statisticRepository) GetAllStatistics(ctx context.Context) (statistics []*models.Statistic, err error) {
	err = r.db.WithContext(ctx).Find(&statistics).Error
	if err != nil {
		return nil, err
	}
	return statistics, nil
}

func (r *statisticRepository) UpsertStatistic(ctx context.Context, statistic *models.Statistic) error {
	var existingStatistic models.Statistic
	result := r.db.WithContext(ctx).Where("endpoint = ? AND user_id = ?", statistic.Endpoint, statistic.UserID).First(&existingStatistic)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Insert new statistic
		statistic.CreatedAt = time.Now()
		result = r.db.Create(statistic)
	} else {
		// Update existing statistic
		existingStatistic.UpdatedAt = statistic.UpdatedAt
		existingStatistic.Count += statistic.Count
		result = r.db.Save(&existingStatistic)
	}

	return result.Error
}

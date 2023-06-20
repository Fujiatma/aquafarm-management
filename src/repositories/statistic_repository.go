package repositories

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
	"gorm.io/gorm"
)

type StatisticRepository struct {
	db *gorm.DB
}

func NewStatisticRepository(db *gorm.DB) *StatisticRepository {
	return &StatisticRepository{db: db}
}

func (r *StatisticRepository) CreateStatistic(ctx context.Context, statistic models.Statistic) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(statistic).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (r *StatisticRepository) GetAllStatistics(ctx context.Context) (statistics []models.Statistic, err error) {
	err = r.db.WithContext(ctx).Find(&statistics).Error
	if err != nil {
		return nil, err
	}

	return statistics, nil
}

func (r *StatisticRepository) IncreaseEndpointCount(ctx context.Context, endpoint string, agent string) error {
	return nil
}

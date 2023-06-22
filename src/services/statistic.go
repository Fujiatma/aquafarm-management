package services

import (
	"context"
	"github.com/delos/aquafarm-management/src/repositories"

	"github.com/delos/aquafarm-management/src/models"
)

type StatisticService interface {
	CreateStatistic(ctx context.Context, payload *models.Statistic) error
	GetStatistics(ctx context.Context) ([]*models.Statistic, error)
}

type statisticService struct {
	statisticRepository repositories.StatisticRepository
}

func NewStatisticService(statisticRepository repositories.StatisticRepository) StatisticService {
	return &statisticService{statisticRepository}
}

package services

import (
	"context"
	"github.com/delos/aquafarm-management/src/repositories"

	"github.com/delos/aquafarm-management/src/models"
)

type StatisticService interface {
	IncreaseEndpointCount(ctx context.Context, endpoint string, userAgent string) error
	GetStatistics(ctx context.Context) (map[string]models.Statistic, error)
}

type statisticService struct {
	statisticRepository repositories.StatisticRepository
}

func NewStatisticService(statisticRepository repositories.StatisticRepository) StatisticService {
	return &statisticService{statisticRepository}
}

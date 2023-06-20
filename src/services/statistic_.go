package services

import (
	"context"

	"github.com/delos/aquafarm-management/src/models"
)

func (s *statisticService) IncreaseEndpointCount(ctx context.Context, endpoint string, userAgent string) error {
	return s.statisticRepository.IncreaseEndpointCount(ctx, endpoint, userAgent)
}

func (s *statisticService) GetStatistics(ctx context.Context) ([]models.Statistic, error) {
	return s.statisticRepository.GetAllStatistics(ctx)
}

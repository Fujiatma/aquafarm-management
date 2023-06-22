package services

import (
	"context"

	"github.com/delos/aquafarm-management/src/models"
)

func (s *statisticService) GetStatistics(ctx context.Context) ([]*models.Statistic, error) {
	statistics, err := s.statisticRepository.GetAllStatistics(ctx)
	if err != nil {
		return nil, err
	}
	return statistics, nil
}

func (s *statisticService) CreateStatistic(ctx context.Context, payload *models.Statistic) error {
	err := s.statisticRepository.UpsertStatistic(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

package mocks

import (
	"context"

	"github.com/delos/aquafarm-management/src/models"
	"github.com/stretchr/testify/mock"
)

type StatisticRepositoryMock struct {
	mock.Mock
}

func (m *StatisticRepositoryMock) CreateStatistic(ctx context.Context, statistic models.Statistic) error {
	args := m.Called(ctx, statistic)
	return args.Error(0)
}

func (m *StatisticRepositoryMock) GetAllStatistics(ctx context.Context) ([]*models.Statistic, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.Statistic), args.Error(1)
}

func (m *StatisticRepositoryMock) UpsertStatistic(ctx context.Context, statistic *models.Statistic) error {
	args := m.Called(ctx, statistic)
	return args.Error(0)
}

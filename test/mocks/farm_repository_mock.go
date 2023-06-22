package mocks

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/stretchr/testify/mock"
)

// Mock FarmRepository
type FarmRepositoryMock struct {
	mock.Mock
}

func (m *FarmRepositoryMock) UpdateFarm(ctx context.Context, farm *models.Farm) error {
	args := m.Called(ctx, farm)
	return args.Error(0)
}

func (m *FarmRepositoryMock) GetFarmByID(ctx context.Context, farmID string) (*models.Farm, error) {
	args := m.Called(ctx, farmID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Farm), args.Error(1)
}

func (m *FarmRepositoryMock) GetFarmsByUserID(ctx context.Context, userID string) ([]*models.Farm, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Farm), args.Error(1)
}

func (m *FarmRepositoryMock) SoftDeleteFarm(ctx context.Context, farm *models.Farm) error {
	args := m.Called(ctx, farm)
	return args.Error(0)
}

func (m *FarmRepositoryMock) CreateFarm(ctx context.Context, farm *models.Farm) error {
	args := m.Called(ctx, farm)
	return args.Error(0)
}

func (m *FarmRepositoryMock) GetFarmByName(ctx context.Context, farmName string) (*models.Farm, error) {
	args := m.Called(ctx, farmName)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Farm), args.Error(1)
}

package mocks

import (
	"context"
	
	"github.com/delos/aquafarm-management/src/models"
	"github.com/stretchr/testify/mock"
)

type PondRepositoryMock struct {
	mock.Mock
}

func (m *PondRepositoryMock) CreatePond(ctx context.Context, pond *models.Pond) error {
	args := m.Called(ctx, pond)
	return args.Error(0)
}

func (m *PondRepositoryMock) UpdatePond(ctx context.Context, pond *models.Pond) error {
	args := m.Called(ctx, pond)
	return args.Error(0)
}

func (m *PondRepositoryMock) GetPondByID(ctx context.Context, pondID string) (*models.Pond, error) {
	args := m.Called(ctx, pondID)
	return args.Get(0).(*models.Pond), args.Error(1)
}

func (m *PondRepositoryMock) GetPondsByUserID(ctx context.Context, userID string) ([]*models.Pond, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*models.Pond), args.Error(1)
}

func (m *PondRepositoryMock) GetPondByName(ctx context.Context, pondName string) (*models.Pond, error) {
	args := m.Called(ctx, pondName)
	return args.Get(0).(*models.Pond), args.Error(1)
}

func (m *PondRepositoryMock) SoftDeletePond(ctx context.Context, pond *models.Pond) error {
	args := m.Called(ctx, pond)
	return args.Error(0)
}

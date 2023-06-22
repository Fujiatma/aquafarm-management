package services

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/delos/aquafarm-management/src/repositories"
)

type PondService interface {
	CreatePond(ctx context.Context, pond *models.Pond) error
	UpdatePond(ctx context.Context, pond *models.Pond) (*models.Pond, error)
	GetPondByID(ctx context.Context, pondID string) (pond *models.Pond, err error)
	GetPondsByUserID(ctx context.Context, userID string) (ponds []*models.Pond, err error)
	GetPondByName(ctx context.Context, pondName string) (pond *models.Pond, err error)
	DeletePond(ctx context.Context, pond *models.Pond, userID string) error
}

type pondService struct {
	pondRepository repositories.PondRepository
	farmRepository repositories.FarmRepository
}

func NewPondService(pondRepository repositories.PondRepository, farmRepository repositories.FarmRepository) PondService {
	return &pondService{pondRepository, farmRepository}
}

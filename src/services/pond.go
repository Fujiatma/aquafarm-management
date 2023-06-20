package services

import (
	"context"
	"github.com/delos/aquafarm-management/src/models"
	"github.com/delos/aquafarm-management/src/repositories"
)

type PondService interface {
	CreatePond(ctx context.Context, pond models.Pond) error
	UpdatePond(ctx context.Context, pond models.Pond) error
	GetPondByID(ctx context.Context, pondID string) (pond *models.Pond, err error)
	GetAllPonds(ctx context.Context) (ponds []*models.Pond, err error)
}

type pondService struct {
	pondRepository repositories.PondRepository
}

func NewPondService(pondRepository repositories.PondRepository) PondService {
	return &pondService{pondRepository}
}

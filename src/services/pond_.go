package services

import (
	"context"

	"github.com/delos/aquafarm-management/src/models"
)

func (s pondService) CreatePond(ctx context.Context, pond models.Pond) error {
	//TODO implement me
	panic("implement me")
}

func (s pondService) UpdatePond(ctx context.Context, pond models.Pond) error {
	//TODO implement me
	panic("implement me")
}

func (s pondService) GetPondByID(ctx context.Context, pondID string) (pond *models.Pond, err error) {
	//TODO implement me
	panic("implement me")
}

func (s pondService) GetAllPonds(ctx context.Context) (ponds []*models.Pond, err error) {
	//TODO implement me
	panic("implement me")
}

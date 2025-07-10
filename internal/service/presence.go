package service

import (
	"context"
	"io"
	"projeto-integrador-mdm/internal/database"
	"projeto-integrador-mdm/internal/domain"
)

type PresenceService interface {
	Create(ctx context.Context, body io.ReadCloser) (*domain.Presence, error)
	// List(ctx context.Context) ([]database.Present, error)
	// Delete(ctx context.Context, numberCard string) (int64, error)
}

type presenceService struct {
	repo *database.Queries
}

func NewPresenceService(queries *database.Queries) *presenceService {
	return &presenceService{repo: queries}
}

func (s *presenceService) Create(ctx context.Context, body io.ReadCloser) (*domain.Presence, error) {
	// s.repo.CreatePresence(ctx, arg database.CreatePresenceParams)
	return nil, nil
}

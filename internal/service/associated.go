package service

import (
	"context"
	"io"
	"projeto-integrador-mdm/internal/database"
	"projeto-integrador-mdm/internal/domain"
	"strconv"
)

type AssociatedService interface {
	Create(ctx context.Context, body io.ReadCloser) (*domain.Associated, error)
	List(ctx context.Context) ([]database.Associated, error)
	Delete(ctx context.Context, numberCard string) (int64, error)
}

type associatedService struct {
	repo *database.Queries
}

func NewAssociatedService(queries *database.Queries) *associatedService {
	return &associatedService{repo: queries}
}

func (s *associatedService) List(ctx context.Context) ([]database.Associated, error) {
	return s.repo.GetAssociated(ctx)
}

func (s *associatedService) Create(ctx context.Context, body io.ReadCloser) (*domain.Associated, error) {
	associatedParams, err := serialization[domain.Associated, database.CreateAssociatedParams](body)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateAssociated(ctx, associatedParams); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *associatedService) Delete(ctx context.Context, numberCard string) (int64, error) {
	numberCardInt, err := strconv.ParseInt(numberCard, 10, 32)
	if err != nil {
		return -1, err
	}

	result, err := s.repo.DeleteAssociatedByNumberCard(ctx, numberCardInt)
	if err != nil {
		return -1, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rows, nil
}

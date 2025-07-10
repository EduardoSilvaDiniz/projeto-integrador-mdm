package service

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"projeto-integrador-mdm/internal/database"
	"projeto-integrador-mdm/internal/domain"
)

type AssociatedService interface {
	Create(ctx context.Context, body io.ReadCloser) (*database.CreateAssociatedParams, error)
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

func (s *associatedService) Create(ctx context.Context, body io.ReadCloser) (*database.CreateAssociatedParams, error) {
	var dto domain.Associated
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := IsValid(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()
	if err := s.repo.CreateAssociated(ctx, params); err != nil {
		return nil, err
	}
	return &params, nil
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

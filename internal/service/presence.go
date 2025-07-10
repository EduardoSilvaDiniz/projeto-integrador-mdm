package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"projeto-integrador-mdm/internal/database"
	"projeto-integrador-mdm/internal/domain"
)

type PresenceService interface {
	Create(ctx context.Context, body io.ReadCloser) (*database.CreatePresenceParams, error)
	List(ctx context.Context) ([]database.Presence, error)
	Delete(ctx context.Context, body io.ReadCloser) (*database.DeletePresenceByCompositeKeyParams, error)
}

type presenceService struct {
	repo *database.Queries
}

func NewPresenceService(queries *database.Queries) *presenceService {
	return &presenceService{repo: queries}
}

func (s *presenceService) List(ctx context.Context) ([]database.Presence, error) {
	return s.repo.GetPresence(ctx)
}

func (s *presenceService) Create(ctx context.Context, body io.ReadCloser) (*database.CreatePresenceParams, error) {
	var dto domain.Presence
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := IsValid(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()

	if err := s.repo.CreatePresence(ctx, params); err != nil {
		return nil, err
	}
	return &params, nil
}

// number_card = ? AND meeting_id = ?
func (s *presenceService) Delete(ctx context.Context, body io.ReadCloser) (*database.DeletePresenceByCompositeKeyParams, error) {
	var dto domain.PresenceCompositeKey
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := IsValid(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()

	result, err := s.repo.DeletePresenceByCompositeKey(ctx, params)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("não foi encontrando registros")
	}

	return &params, nil
}

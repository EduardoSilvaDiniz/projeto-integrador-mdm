package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/domain"
)

type PresenceService interface {
	Create(ctx context.Context, body io.ReadCloser) (*db.CreatePresenceParams, error)
	GetById(ctx context.Context, body io.ReadCloser) (*db.Presence, error)
	Update(ctx context.Context, body io.ReadCloser) (*db.UpdatePresenceParams, error)
	List(ctx context.Context) ([]db.Presence, error)
	Delete(ctx context.Context, body io.ReadCloser) (*db.DeletePresenceByCompositeKeyParams, error)
}

type presenceService struct {
	repo *db.Queries
}

func NewPresenceService(queries *db.Queries) *presenceService {
	return &presenceService{repo: queries}
}

func (s *presenceService) List(ctx context.Context) ([]db.Presence, error) {
	return s.repo.GetPresence(ctx)
}

func (s *presenceService) Create(
	ctx context.Context,
	body io.ReadCloser,
) (*db.CreatePresenceParams, error) {
	var dto domain.Presence
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()

	if err := s.repo.CreatePresence(ctx, params); err != nil {
		return nil, err
	}
	return &params, nil
}

func (s *presenceService) GetById(ctx context.Context, body io.ReadCloser) (*db.Presence, error) {
	var dto domain.PresenceByCompositeKey
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := IsValid(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()

	register, err := s.repo.GetPresenceByCompositeKey(
		ctx,
		db.GetPresenceByCompositeKeyParams(params),
	)
	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (s *presenceService) Update(
	ctx context.Context,
	body io.ReadCloser,
) (*db.UpdatePresenceParams, error) {
	var dto domain.Presence

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	params := dto.ToUpdateParams()

	result, err := s.repo.UpdatePresence(ctx, params)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("not found")
	}

	return &params, nil
}

// number_card = ? AND meeting_id = ?
func (s *presenceService) Delete(
	ctx context.Context,
	body io.ReadCloser,
) (*db.DeletePresenceByCompositeKeyParams, error) {
	var dto domain.PresenceByCompositeKey
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

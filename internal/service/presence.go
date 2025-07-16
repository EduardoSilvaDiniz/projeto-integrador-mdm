package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/domain"
	"projeto-integrador-mdm/internal/errs"
	"strconv"
	"strings"
)

type PresenceService interface {
	Create(ctx context.Context, body io.ReadCloser) (*db.CreatePresenceParams, error)
	GetById(ctx context.Context, numberCard string, meetingId string) (*db.Presence, error)
	Update(ctx context.Context, body io.ReadCloser) (*db.UpdatePresenceParams, error)
	List(ctx context.Context) ([]db.Presence, error)
	Delete(ctx context.Context, body io.ReadCloser) (*db.DeletePresenceByCompositeKeyParams, error)
}

type presenceService struct {
	repo *db.Queries
}

func NewPresenceService(queries *db.Queries) *presenceService {
	defer slog.Debug("objeto presenceService criado")
	return &presenceService{repo: queries}
}

func (s *presenceService) List(ctx context.Context) ([]db.Presence, error) {
	slog.Debug("chamada de sistema", "func", "presenceService.List")
	return s.repo.GetPresence(ctx)
}

func (s *presenceService) Create(
	ctx context.Context,
	body io.ReadCloser,
) (*db.CreatePresenceParams, error) {
	slog.Debug("chamada de sistema", "func", "presenceService.Create")
	var dto domain.Presence
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()
	if err := s.repo.CreatePresence(ctx, params); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("%w: %s", errs.ErrAlreadyExists, err.Error())
		}
		return nil, err
	}
	return &params, nil
}

// TODO o presence getbyid não aceita id, precisa enviar um body contento meetingid e associatedid
func (s *presenceService) GetById(
	ctx context.Context,
	numberCard string,
	meetingId string,
) (*db.Presence, error) {
	slog.Debug("chamada de sistema", "func", "presenceService.GetById")

	numberCardInt, err := strconv.ParseInt(numberCard, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		slog.Error("foi passado algo diferente de int")
		return nil, fmt.Errorf("%w %s", errs.ErrInvalidInput, "só pode ser passado numeros")
	}

	meetingIdInt, err := strconv.ParseInt(meetingId, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		slog.Error("foi passado algo diferente de int")
		return nil, fmt.Errorf("%w %s", errs.ErrInvalidInput, "só pode ser passado numeros")
	}

	presenceByCompositeKey := db.GetPresenceByCompositeKeyParams{
		NumberCard: numberCardInt,
		MeetingID:  meetingIdInt,
	}
	register, err := s.repo.GetPresenceByCompositeKey(ctx, presenceByCompositeKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		slog.Error("erro de execução", "func", "repo.GetAssociatedByNumberCard")
		return nil, err
	}

	return &register, nil
}

func (s *presenceService) Update(
	ctx context.Context,
	body io.ReadCloser,
) (*db.UpdatePresenceParams, error) {
	slog.Debug("chamada de sistema", "func", "presenceService.Update")

	data, err := decodeJson(body)
	if err != nil {
		slog.Error("erro de execução", "func", "decodeJson")
		return nil, err
	}

	var dto domain.Presence

	if err = json.Unmarshal(data, &dto); err != nil {
		slog.Error("erro de execução", "func", "json.Unmarshal")
		return nil, err
	}

	if err = ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToUpdateParams()

	result, err := s.repo.UpdatePresence(ctx, params)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.UpdatePresence")
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		slog.Error("erro de execução", "func", "result.RowsAffected")
		return nil, err
	}

	if rows == 0 {
		return nil, nil
	}

	return &params, nil
}

func (s *presenceService) Delete(
	ctx context.Context,
	body io.ReadCloser,
) (*db.DeletePresenceByCompositeKeyParams, error) {
	slog.Debug("chamada de função Delete do presenceService")
	var dto domain.PresenceByCompositeKey
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToDeleteParams()

	result, err := s.repo.DeletePresenceByCompositeKey(ctx, params)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.DeletePresenceByCompositeKey")
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		slog.Error("erro de execução", "func", "result.RowsAffected")
		return nil, err
	}

	if rows == 0 {
		return nil, nil
	}

	return &params, nil
}

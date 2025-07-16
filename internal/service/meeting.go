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

type MeetingService interface {
	Create(ctx context.Context, body io.ReadCloser) (*db.CreateMeetingParams, error)
	GetById(ctx context.Context, id string) (*db.Meeting, error)
	Update(ctx context.Context, body io.ReadCloser) (*db.UpdateMeetingParams, error)
	List(ctx context.Context) ([]db.Meeting, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type meetingService struct {
	repo *db.Queries
}

func NewMeetingService(queries *db.Queries) *meetingService {
	defer slog.Debug("objeto meetingService criado")
	return &meetingService{repo: queries}
}

func (s *meetingService) List(ctx context.Context) ([]db.Meeting, error) {
	slog.Debug("chamada de sistema", "func", "meetingService.List")
	return s.repo.GetMeetings(ctx)
}

func (s *meetingService) Create(
	ctx context.Context,
	body io.ReadCloser,
) (*db.CreateMeetingParams, error) {
	slog.Debug("chamada de sistema", "func", "meetingService.Create")
	var dto domain.Meeting
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()
	if err := s.repo.CreateMeeting(ctx, params); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("%w: %s", errs.ErrAlreadyExists, err.Error())
		}
		return nil, err
	}
	return &params, nil
}

func (s *meetingService) GetById(ctx context.Context, id string) (*db.Meeting, error) {
	slog.Debug("chamada de sistema", "func", "meetingService.GetById")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		slog.Error("foi passado algo diferente de int")
		return nil, fmt.Errorf("%w %s", errs.ErrInvalidInput, "só pode ser passado numeros")
	}

	register, err := s.repo.GetMeetingById(ctx, idInt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		slog.Error("erro de execução", "func", "repo.GetMeetingById")
		return nil, err
	}

	return &register, nil
}

func (s *meetingService) Update(
	ctx context.Context,
	body io.ReadCloser,
) (*db.UpdateMeetingParams, error) {
	slog.Debug("chamada de sistema", "func", "meetingService.Update")

	data, err := decodeJson(body)
	if err != nil {
		slog.Error("erro de execução", "func", "decodeJson")
		return nil, err
	}

	var dto domain.Meeting

	if err = json.Unmarshal(data, &dto); err != nil {
		slog.Error("erro de execução", "func", "json.Unmarshal")
		return nil, err
	}

	if err = ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToUpdateParams()

	result, err := s.repo.UpdateMeeting(ctx, params)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.UpdateMeeting")
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

func (s *meetingService) Delete(ctx context.Context, id string) (int64, error) {
	slog.Debug("chamada de função Delete do meetingService")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		return -1, err
	}

	result, err := s.repo.DeleteMeetingById(ctx, idInt)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.DeleteMeetingById")
		return -1, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		slog.Error("erro de execução", "func", "result.RowsAffected")
		return -1, err
	}

	if rows == 0 {
		return 0, nil
	}

	return rows, nil
}

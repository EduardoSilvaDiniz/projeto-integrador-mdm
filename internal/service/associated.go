package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log/slog"
	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/domain"
	"strconv"
)

type AssociatedService interface {
	Create(ctx context.Context, body io.ReadCloser) (*db.CreateAssociatedParams, error)
	GetById(ctx context.Context, id string) (*db.Associated, error)
	Update(ctx context.Context, body io.ReadCloser) (*db.UpdateAssociatedParams, error)
	List(ctx context.Context) ([]db.Associated, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type associatedService struct {
	repo *db.Queries
}

func NewAssociatedService(queries *db.Queries) *associatedService {
	defer slog.Debug("objeto associatedService criado")
	return &associatedService{repo: queries}
}

func (s *associatedService) List(ctx context.Context) ([]db.Associated, error) {
	slog.Debug("chamada de sistema", "func", "associatedService.List")
	return s.repo.GetAssociated(ctx)
}

func (s *associatedService) Create(
	ctx context.Context,
	body io.ReadCloser,
) (*db.CreateAssociatedParams, error) {
	slog.Debug("chamada de sistema", "func", "associatedService.Create")
	var dto domain.Associated
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()
	if err := s.repo.CreateAssociated(ctx, params); err != nil {
		return nil, err
	}
	return &params, nil
}

func (s *associatedService) GetById(ctx context.Context, id string) (*db.Associated, error) {
	slog.Debug("chamada de sistema", "func", "associatedService.GetById")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		return nil, err
	}

	register, err := s.repo.GetAssociatedByNumberCard(ctx, idInt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		slog.Error("erro de execução", "func", "repo.GetAssociatedByNumberCard")
		return nil, err
	}

	return &register, nil
}

func (s *associatedService) Update(
	ctx context.Context,
	body io.ReadCloser,
) (*db.UpdateAssociatedParams, error) {
	slog.Debug("chamada de sistema", "func", "associatedService.Update")

	data, err := decodeJson(body)
	if err != nil {
		slog.Error("erro de execução", "func", "decodeJson")
		return nil, err
	}

	var dto domain.Associated

	if err = json.Unmarshal(data, &dto); err != nil {
		slog.Error("erro de execução", "func", "json.Unmarshal")
		return nil, err
	}

	params := dto.ToUpdateParams()

	result, err := s.repo.UpdateAssociated(ctx, params)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.UpdateAssociated")
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

func (s *associatedService) Delete(ctx context.Context, id string) (int64, error) {
	slog.Debug("chamada de função Delete do associatedService")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		return -1, err
	}

	result, err := s.repo.DeleteAssociatedByNumberCard(ctx, idInt)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.DeleteAssociatedByNumberCard")
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

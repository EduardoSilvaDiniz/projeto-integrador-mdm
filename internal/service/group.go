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

type GroupService interface {
	Create(ctx context.Context, body io.ReadCloser) (*db.CreateGroupParams, error)
	GetById(ctx context.Context, id string) (*db.Group, error)
	Update(ctx context.Context, body io.ReadCloser) (*db.UpdateGroupParams, error)
	List(ctx context.Context) ([]db.Group, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type groupService struct {
	repo *db.Queries
}

func NewGroupService(queries *db.Queries) *groupService {
	defer slog.Debug("objeto groupService criado")
	return &groupService{repo: queries}
}

func (s *groupService) List(ctx context.Context) ([]db.Group, error) {
	slog.Debug("chamada de sistema", "func", "groupService.List")
	return s.repo.GetGroups(ctx)
}

func (s *groupService) Create(
	ctx context.Context,
	body io.ReadCloser,
) (*db.CreateGroupParams, error) {
	slog.Debug("chamada de sistema", "func", "groupService.Create")
	var dto domain.Group
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()
	if err := s.repo.CreateGroup(ctx, params); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("%w: %s", errs.ErrAlreadyExists, err.Error())
		}
		return nil, err
	}
	return &params, nil
}

func (s *groupService) GetById(ctx context.Context, id string) (*db.Group, error) {
	slog.Debug("chamada de sistema", "func", "groupService.GetById")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		slog.Error("foi passado algo diferente de int")
		return nil, fmt.Errorf("%w %s", errs.ErrInvalidInput, "só pode ser passado numeros")
	}

	register, err := s.repo.GetGroupById(ctx, idInt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		slog.Error("erro de execução", "func", "repo.GetGroupById")
		return nil, err
	}

	return &register, nil
}

func (s *groupService) Update(
	ctx context.Context,
	body io.ReadCloser,
) (*db.UpdateGroupParams, error) {
	slog.Debug("chamada de sistema", "func", "groupService.Update")

	data, err := decodeJson(body)
	if err != nil {
		slog.Error("erro de execução", "func", "decodeJson")
		return nil, err
	}

	var dto domain.Group

	if err = json.Unmarshal(data, &dto); err != nil {
		slog.Error("erro de execução", "func", "json.Unmarshal")
		return nil, err
	}

	if err = ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToUpdateParams()

	result, err := s.repo.UpdateGroup(ctx, params)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.UpdateGroup")
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

func (s *groupService) Delete(ctx context.Context, id string) (int64, error) {
	slog.Debug("chamada de função Delete do groupService")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		return -1, err
	}

	result, err := s.repo.DeleteGroupById(ctx, idInt)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.DeleteGroupById")
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

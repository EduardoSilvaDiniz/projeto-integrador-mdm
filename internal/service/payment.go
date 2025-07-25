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

type PaymentService interface {
	Create(ctx context.Context, body io.ReadCloser) (*db.CreatePaymentParams, error)
	GetById(ctx context.Context, id string) (*db.Payment, error)
	Update(ctx context.Context, body io.ReadCloser) (*db.UpdatePaymentParams, error)
	List(ctx context.Context) ([]db.Payment, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type paymentService struct {
	repo *db.Queries
}

func NewPaymentService(queries *db.Queries) *paymentService {
	defer slog.Debug("objeto paymentService criado")
	return &paymentService{repo: queries}
}

func (s *paymentService) List(ctx context.Context) ([]db.Payment, error) {
	slog.Debug("chamada de sistema", "func", "paymentService.List")
	return s.repo.GetPayment(ctx)
}

func (s *paymentService) Create(
	ctx context.Context,
	body io.ReadCloser,
) (*db.CreatePaymentParams, error) {
	slog.Debug("chamada de sistema", "func", "paymentService.Create")
	var dto domain.Payment
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()
	if err := s.repo.CreatePayment(ctx, params); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("%w: %s", errs.ErrAlreadyExists, err.Error())
		}
		return nil, err
	}
	return &params, nil
}

func (s *paymentService) GetById(ctx context.Context, id string) (*db.Payment, error) {
	slog.Debug("chamada de sistema", "func", "paymentService.GetById")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		slog.Error("foi passado algo diferente de int")
		return nil, fmt.Errorf("%w %s", errs.ErrInvalidInput, "só pode ser passado numeros")
	}

	register, err := s.repo.GetPaymentById(ctx, idInt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		slog.Error("erro de execução", "func", "repo.GetPaymentById")
		return nil, err
	}

	return &register, nil
}

func (s *paymentService) Update(
	ctx context.Context,
	body io.ReadCloser,
) (*db.UpdatePaymentParams, error) {
	slog.Debug("chamada de sistema", "func", "paymentService.Update")

	data, err := decodeJson(body)
	if err != nil {
		slog.Error("erro de execução", "func", "decodeJson")
		return nil, err
	}

	var dto domain.Payment

	if err = json.Unmarshal(data, &dto); err != nil {
		slog.Error("erro de execução", "func", "json.Unmarshal")
		return nil, err
	}

	if err = ValidateStruct(dto); err != nil {
		return nil, err
	}

	params := dto.ToUpdateParams()

	result, err := s.repo.UpdatePayment(ctx, params)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.UpdatePayment")
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

func (s *paymentService) Delete(ctx context.Context, id string) (int64, error) {
	slog.Debug("chamada de função Delete do paymentService")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		slog.Error("erro de execução", "func", "strconv.ParseInt")
		return -1, err
	}

	result, err := s.repo.DeletePaymentById(ctx, idInt)
	if err != nil {
		slog.Error("erro de execução", "func", "repo.DeletePaymentById")
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

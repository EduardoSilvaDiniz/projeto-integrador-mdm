package service

import (
	"context"
	"encoding/json"
	"io"
	"projeto-integrador-mdm/internal/database"
	"projeto-integrador-mdm/internal/domain"
	"strconv"
)

type PaymentService interface {
	Create(ctx context.Context, body io.ReadCloser) (*database.CreatePaymentParams, error)
	List(ctx context.Context) ([]database.Payment, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type paymentService struct {
	repo *database.Queries
}

func NewPaymentService(queries *database.Queries) *paymentService {
	return &paymentService{repo: queries}
}

func (s *paymentService) List(ctx context.Context) ([]database.Payment, error) {
	return s.repo.GetPayment(ctx)
}

func (s *paymentService) Create(ctx context.Context, body io.ReadCloser) (*database.CreatePaymentParams, error) {
	var dto domain.Payment
	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	if err := IsValid(dto); err != nil {
		return nil, err
	}

	params := dto.ToCreateParams()
	if err := s.repo.CreatePayment(ctx, params); err != nil {
		return nil, err
	}
	return &params, nil
}

func (s *paymentService) Delete(ctx context.Context, id string) (int64, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return -1, err
	}

	result, err := s.repo.DeletePaymentById(ctx, idInt)
	if err != nil {
		return -1, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rows, nil
}

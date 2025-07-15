package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/domain"
	"strconv"
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
	return &paymentService{repo: queries}
}

func (s *paymentService) List(ctx context.Context) ([]db.Payment, error) {
	return s.repo.GetPayment(ctx)
}

func (s *paymentService) Create(ctx context.Context, body io.ReadCloser) (*db.CreatePaymentParams, error) {
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

func (s *paymentService) GetById(ctx context.Context, id string) (*db.Payment, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, err
	}

	register, err := s.repo.GetPaymentByID(ctx, idInt)
	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (s *paymentService) Update(ctx context.Context, body io.ReadCloser) (*db.UpdatePaymentParams, error) {
	var dto domain.Payment

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	params := dto.ToUpdateParams()

	result, err := s.repo.UpdatePayment(ctx, params)
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

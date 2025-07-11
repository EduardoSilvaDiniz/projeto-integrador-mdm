package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"strconv"

	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/domain"
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
	return &associatedService{repo: queries}
}

func (s *associatedService) List(ctx context.Context) ([]db.Associated, error) {
	return s.repo.GetAssociated(ctx)
}

func (s *associatedService) Create(ctx context.Context, body io.ReadCloser) (*db.CreateAssociatedParams, error) {
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

func (s *associatedService) GetById(ctx context.Context, id string) (*db.Associated, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, err
	}

	register, err := s.repo.GetAssociatedByNumberCard(ctx, idInt)
	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (s *associatedService) Update(ctx context.Context, body io.ReadCloser) (*db.UpdateAssociatedParams, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		log.Println("erro ao ler body:", err)
		return nil, err
	}
	log.Println("corpo recebido:", string(data))

	var dto domain.Associated

	if err := json.Unmarshal(data, &dto); err != nil {
		log.Println("erro de decodificação Json")
		return nil, err
	}

	params := dto.ToUpdateParams()

	result, err := s.repo.UpdateAssociated(ctx, params)
	if err != nil {
		log.Println("erro de execução SQL")
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("erro de execução RowsAffected")
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("not found")
	}

	return &params, nil
}

func (s *associatedService) Delete(ctx context.Context, id string) (int64, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return -1, err
	}

	result, err := s.repo.DeleteAssociatedByNumberCard(ctx, idInt)
	if err != nil {
		return -1, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rows, nil
}

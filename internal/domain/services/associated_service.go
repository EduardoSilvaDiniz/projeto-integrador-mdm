package services

import (
	"chamada-pagamento-system/internal/infra/repositories"
	"chamada-pagamento-system/internal/transport/http-server/dto"
)

type AssociatedService struct {
	Repo *repositories.GormAssociatedRepository
}

func NewAssociatedService(repo *repositories.GormAssociatedRepository) *AssociatedService {
	return &AssociatedService{Repo: repo}
}

func (s *AssociatedService) Create(a *dto.Associated) error {
	if err := a.IsValid(); err != nil {
		return err
	}

	if err := s.Repo.Save(a.ToEntity()); err != nil {
		return err
	}

	return nil
}

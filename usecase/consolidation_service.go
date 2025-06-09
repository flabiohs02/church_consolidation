package usecase

import (
	"church_consolidation/domain"
	"church_consolidation/repository"
)

type ConsolidationService struct {
	repo repository.ConsolidationRepository
}

func NewConsolidationService(repo repository.ConsolidationRepository) *ConsolidationService {
	return &ConsolidationService{repo: repo}
}

func (s *ConsolidationService) CreateConsolidation(consolidation *domain.Consolidation) error {
	return s.repo.Save(consolidation)
}

func (s *ConsolidationService) GetConsolidationByID(id uint) (*domain.Consolidation, error) {
	return s.repo.FindByID(id)
}

func (s *ConsolidationService) GetAllConsolidations() ([]domain.Consolidation, error) {
	return s.repo.FindAll()
}

func (s *ConsolidationService) UpdateConsolidation(consolidation *domain.Consolidation) error {
	return s.repo.Update(consolidation)
}

func (s *ConsolidationService) DeleteConsolidation(id uint) error {
	return s.repo.Delete(id)
} 
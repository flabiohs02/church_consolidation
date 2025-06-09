package repository

import (
	"church_consolidation/domain"
)

type ConsolidationRepository interface {
	Save(consolidation *domain.Consolidation) error
	FindByID(id uint) (*domain.Consolidation, error)
	FindAll() ([]domain.Consolidation, error)
	Update(consolidation *domain.Consolidation) error
	Delete(id uint) error
}

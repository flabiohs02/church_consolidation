package repository

import (
	"church_consolidation/domain"

	"gorm.io/gorm"
)

type GormConsolidationRepository struct {
	db *gorm.DB
}

func NewGormConsolidationRepository(db *gorm.DB) *GormConsolidationRepository {
	return &GormConsolidationRepository{db: db}
}

func (r *GormConsolidationRepository) Save(consolidation *domain.Consolidation) error {
	return r.db.Create(consolidation).Error
}

func (r *GormConsolidationRepository) FindByID(id uint) (*domain.Consolidation, error) {
	var consolidation domain.Consolidation
	err := r.db.First(&consolidation, id).Error
	return &consolidation, err
}

func (r *GormConsolidationRepository) FindAll() ([]domain.Consolidation, error) {
	var consolidations []domain.Consolidation
	err := r.db.Find(&consolidations).Error
	return consolidations, err
}

func (r *GormConsolidationRepository) Update(consolidation *domain.Consolidation) error {
	return r.db.Save(consolidation).Error
}

func (r *GormConsolidationRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Consolidation{}, id).Error
}

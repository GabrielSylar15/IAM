package repository

import (
	"IAM/entities"
	"github.com/jinzhu/gorm"
)

type scopeRepository struct {
	db *gorm.DB
}

type ScopeRepository interface {
	CreateScope(scope *entities.Scope) (entities.Scope, error)
}

func InitScopeRepository(db *gorm.DB) ScopeRepository {
	return &scopeRepository{db}
}

func (r *scopeRepository) CreateScope(scope *entities.Scope) (entities.Scope, error) {
	err := r.db.Create(scope).Error
	return *scope, err
}

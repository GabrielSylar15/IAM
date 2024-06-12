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
	GetScope(clientID string) ([]*entities.Scope, error)
}

func InitScopeRepository(db *gorm.DB) ScopeRepository {
	return &scopeRepository{db}
}

func (r *scopeRepository) CreateScope(scope *entities.Scope) (entities.Scope, error) {
	err := r.db.Create(scope).Error
	return *scope, err
}

func (r *scopeRepository) GetScope(clientID string) ([]*entities.Scope, error) {
	var scopes []*entities.Scope
	// a = &b => *a = b;
	err := r.db.Where("owner_client = ?", clientID).Find(&scopes).Error
	return scopes, err
}

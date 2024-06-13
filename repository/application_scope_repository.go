package repository

import (
	"IAM/entities"
	"github.com/jinzhu/gorm"
)

type applicationScopeRepository struct {
	DB *gorm.DB
}

type ApplicationScopeRepository interface {
	AssignScope(entity *entities.ApplicationScope) (*entities.ApplicationScope, error)
}

func InitApplicationScopeRepository(db *gorm.DB) ApplicationScopeRepository {
	return &applicationScopeRepository{db}
}

func (r *applicationScopeRepository) AssignScope(entity *entities.ApplicationScope) (*entities.ApplicationScope, error) {
	err := r.DB.Create(entity).Error
	return entity, err
}

package repository

import (
	"IAM/entities"
	"github.com/jinzhu/gorm"
)

type ApplicationRepository interface {
	CreateApplication(service *entities.Application) error
}

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepository{db}
}

func (r *applicationRepository) CreateApplication(service *entities.Application) error {
	return r.db.Create(service).Error
}

package repository

import (
	"IAM/entities"
	"github.com/jinzhu/gorm"
)

type ApplicationRepository interface {
	CreateApplication(service *entities.Application) error
}

// định nghĩa class impl interface
type applicationRepository struct {
	db *gorm.DB
}

// nó giống một constructor
// define trả về interface nhưng return địa chỉ của struct => interface trỏ tới địa chỉ của class
func InitApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepository{db}
}

// receiver(method) khác với function
// nó có thể trùng tên, và nó cho biết nó thuộc struct(class) nào
func (r *applicationRepository) CreateApplication(service *entities.Application) error {
	return r.db.Create(service).Error
}

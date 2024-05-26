package entities

import (
	"time"
)

type Application struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement;comment:'id service_scope'"`
	ClientID     string    `gorm:"column:client_id;type:varchar(255);not null;comment:'id client service'"`
	ClientSecret string    `gorm:"column:client_secret;type:varchar(255);not null"`
	ServiceName  string    `gorm:"column:service_name;type:varchar(100);not null"`
	PublicKey    string    `gorm:"column:public_key;type:varchar(2000);not null"`
	PrivateKey   string    `gorm:"column:private_key;type:varchar(2000);not null"`
	CreatedBy    uint64    `gorm:"column:created_by;not null"`
	KeyID        string    `gorm:"column:kid;type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:''"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:''"`
}

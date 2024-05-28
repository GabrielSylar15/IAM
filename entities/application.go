package entities

import (
	"time"
)

type Application struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement;comment:'id service_scope'" json:"id"`
	ClientID     string    `gorm:"column:client_id;type:varchar(255);not null;comment:'id client service'" json:"client_id"`
	ClientSecret string    `gorm:"column:client_secret;type:varchar(255);not null" json:"client_secret"`
	ServiceName  string    `gorm:"column:service_name;type:varchar(100);not null" json:"service_name"`
	PublicKey    string    `gorm:"column:public_key;type:varchar(2000);not null" json:"public_key"`
	PrivateKey   string    `gorm:"column:private_key;type:varchar(2000);not null" json:"private_key"`
	CreatedBy    uint64    `gorm:"column:created_by;not null" json:"created_by"`
	KeyID        string    `gorm:"column:kid;type:varchar(255);not null" json:"key_id"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:''" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:''" json:"updated_at"`
}

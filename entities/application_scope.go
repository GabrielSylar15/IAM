package entities

import (
	"time"
)

type ApplicationScope struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;comment:'id service_scope'" json:"id"`
	ServiceID uint64    `gorm:"column:service_id;not null;comment:'id của service'" json:"service_id"`
	ScopeID   uint64    `gorm:"column:scope_id;not null;comment:'id của scope'" json:"scope_id"`
	ClientID  string    `gorm:"column:client_id;type:varchar(255);not null;comment:'id client service'" json:"client_id"`
	CreatedBy uint64    `gorm:"column:created_by;not null;comment:'id của người thêm scope'" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:''" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:''" json:"updated_at"`
}

func (ApplicationScope) TableName() string {
	return "service_scope"
}

package entities

import (
	"time"
)

type ApplicationScope struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;comment:'id service_scope'"`
	ServiceID uint64    `gorm:"column:service_id;not null;comment:'id của service'"`
	ScopeID   uint64    `gorm:"column:scope_id;not null;comment:'id của scope'"`
	ClientID  string    `gorm:"column:client_id;type:varchar(255);not null;comment:'id client service'"`
	CreatedBy uint64    `gorm:"column:created_by;not null;comment:'id của người thêm scope'"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:''"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:''"`
}

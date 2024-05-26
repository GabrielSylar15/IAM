package entities

import (
	"time"
)

type Scope struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Scp         string    `gorm:"column:scp;type:varchar(255);not null"`
	OwnerClient string    `gorm:"column:owner_client;type:varchar(255);not null"`
	CreatedBy   uint64    `gorm:"column:created_by;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:''"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:''"`
}

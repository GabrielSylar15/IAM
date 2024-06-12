package entities

import (
	"time"
)

type Scope struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Scp         string    `gorm:"column:scp;type:varchar(255);not null" json:"scp"`
	OwnerClient string    `gorm:"column:owner_client;type:varchar(255);not null" json:"owner_client"`
	CreatedBy   uint64    `gorm:"column:created_by;not null" json:"created_by"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:''" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:''" json:"updated_at"`
}

func (Scope) TableName() string {
	return "scope"
}

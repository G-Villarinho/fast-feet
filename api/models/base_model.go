package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt sql.NullTime   `gorm:"default:null"`
	DeletedAt gorm.DeletedAt `gorm:"default:null;index"`
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	return
}

package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	UserName  string    `gorm:"not null"`
	Farms     []Farm    `gorm:"-"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

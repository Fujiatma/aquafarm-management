package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	UserName  string    `gorm:"not null"`
	Farms     []Farm    `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

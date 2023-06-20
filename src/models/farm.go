package models

import "time"

type Farm struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	FarmName  string `gorm:"column:farm_name"`
	UserID    string
	IsDeleted bool      `gorm:"column:is_deleted"`
	User      User      `gorm:"foreignKey:UserID"`
	Ponds     []Pond    `gorm:"foreignKey:FarmID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt time.Time
}

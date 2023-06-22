package models

import "time"

type Farm struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	FarmName  string    `gorm:"column:farm_name"`
	UserID    string    `gorm:"column:user_id"`
	IsDeleted bool      `gorm:"column:is_deleted"`
	Ponds     []*Pond   `gorm:"foreignKey:FarmID"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time
}

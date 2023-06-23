package models

import "time"

type Pond struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	PondName  string `gorm:"column:pond_name"`
	FarmID    string
	IsDeleted bool      `gorm:"column:is_deleted"`
	Farm      Farm      `gorm:"foreignKey:FarmID"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time
}

package models

import "time"

type Statistic struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	Endpoint  string    `gorm:"column:endpoint"`
	UserID    string    `gorm:"type:varchar(36);column:user_id"`
	Count     int       `gorm:"column:count"`
	CallAt    time.Time `gorm:"column:call_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

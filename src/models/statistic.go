package models

import "time"

type Statistic struct {
	ID              string    `gorm:"type:varchar(36);primaryKey"`
	Endpoint        string    `gorm:"column:endpoint"`
	Count           int       `gorm:"column:count"`
	UniqueUserAgent int       `gorm:"column:unique_user_agent"`
	UserAgent       string    `gorm:"column:user_agent"`
	CallAt          time.Time `gorm:"column:call_at"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoCreateTime"`
}

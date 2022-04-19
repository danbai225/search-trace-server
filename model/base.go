package model

import "time"

type Base struct {
	CreatedAt *time.Time `gorm:"column:created_at;comment:'创建时间'" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;comment:'更新时间'" json:"updated_at"`
}

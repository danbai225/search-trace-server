package model

import "time"

type Base struct {
	Id        int64      `gorm:"column:id;type:BIGINT;autoIncrement;comment:'唯一id';NOT NULL" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at;comment:'创建时间'" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;comment:'更新时间'" json:"updated_at"`
}

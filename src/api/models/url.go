package models

import (
	"time"
)

type Site struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	URL         string    `gorm:"column:url;not null;type:varchar(255)" json:"url"`
	Domain      string    `gorm:"column:domain;not null;type:varchar(100)" json:"domain"`
	Description string    `gorm:"column:description;type:varchar(255)" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

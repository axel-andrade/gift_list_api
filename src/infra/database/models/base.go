package models

import (
	"time"
)

type Base struct {
	ID        int64     `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}

package entity

import "gorm.io/gorm"

type ApexTracker struct {
	gorm.Model
	content string `gorm:"not null"`
}

package entity

import "gorm.io/gorm"

type ApexTracker struct {
	gorm.Model
	Content string `gorm:"not null"`
}

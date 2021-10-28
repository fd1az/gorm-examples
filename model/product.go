package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string `gorm:"uniqueIndex"`
	Description string
}

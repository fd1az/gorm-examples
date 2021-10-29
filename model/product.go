package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `gorm:"uniqueIndex"`
	Description string
	Price       decimal.Decimal
}

func (Product) TableName() string {
	return "public.products"
}

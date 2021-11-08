package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User         User
	UserID       uint
	Total        decimal.Decimal
	Status       string
	OrderProduct []OrderProduct
}

func (Order) TableName() string {
	return "public.orders"
}

type OrderProduct struct {
	OrderID   uint
	ProductID uint
	UnitPrice decimal.Decimal
	Quantity  int
}

func (OrderProduct) TableName() string {
	return "public.order_products"
}

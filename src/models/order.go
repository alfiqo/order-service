package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID          uint `json:"id,omitempty"`
	OrderNumber string
	MerchantId  string  `json:"merchantId" binding:"required"`
	ProductId   string  `json:"productId" binding:"required"`
	QTY         uint    `json:"qty" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Discount    float64 `json:"discount"`
	Status      string  `json:"status"`
}

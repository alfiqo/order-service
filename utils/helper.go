package utils

import (
	"math"
	"order-service/src/responses"
	"time"

	"gorm.io/gorm"
)

func ValidateDateFormat(date string) bool {
	_, err := time.Parse(time.DateOnly, date)
	return err == nil
}

func Paginate(value interface{}, pagination *responses.Meta, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalData = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetPerPage())))
	pagination.MaxPage = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPerPage())
	}
}

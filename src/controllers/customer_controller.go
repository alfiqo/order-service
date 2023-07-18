package controllers

import (
	"net/http"
	"order-service/config"
	"order-service/src/models"
	"order-service/src/resonses/customers"
	"order-service/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.NewDB()

func Create(context *gin.Context) {
	var customer models.Customer

	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customer.DOB != nil {
		ok := utils.ValidateDateFormat(customer.DOB)
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{"error": utils.DateFormatInvalid})
			return
		}
	}

	exist := db.Where("email = ?", customer.Email).First(&customer)
	if exist.Error == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": utils.EmailRegistered})
		return
	}

	result := db.Create(&customer)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	context.JSON(http.StatusOK, customers.NewCreateResponse(&customer))
}

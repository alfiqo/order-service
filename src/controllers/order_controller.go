package controllers

import (
	"order-service/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController() OrderController {
	return OrderController{
		DB: config.NewDB(),
	}
}

func (c *OrderController) GetAll(context *gin.Context) {}
func (c *OrderController) Create(context *gin.Context) {}
func (c *OrderController) Update(context *gin.Context) {}
func (c *OrderController) Detail(context *gin.Context) {}
func (c *OrderController) Delete(context *gin.Context) {}

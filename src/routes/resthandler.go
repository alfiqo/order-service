package routes

import (
	"net/http"
	"order-service/src/controllers"

	"github.com/gin-gonic/gin"
)

func NewRoutes() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := route.Group("/api/v1")

	v1.POST("customers", controllers.Create)

	route.Run()
}

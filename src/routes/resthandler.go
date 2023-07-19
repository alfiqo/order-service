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

	v1.GET("customers", controllers.GetAll)
	v1.POST("customers", controllers.Create)
	v1.GET("customers/:id", controllers.Detail)
	v1.PUT("customers/:id", controllers.Update)
	v1.DELETE("customers/:id", controllers.Delete)

	v1.POST("/register", controllers.Register)

	route.Run()
}

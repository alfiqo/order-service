package routes

import (
	"net/http"
	"order-service/src/controllers"
	"order-service/src/middleware"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userController := controllers.NewUserController()
	route.POST("/register", userController.Register)
	route.POST("/login", userController.Login)

	v1 := route.Group("/api/v1")

	v1.Use(middleware.JwtAuthMiddleware())

	customerController := controllers.NewCustomerController()

	v1.GET("customers", customerController.GetAll)
	v1.POST("customers", customerController.Create)
	v1.GET("customers/:id", customerController.Detail)
	v1.PUT("customers/:id", customerController.Update)
	v1.DELETE("customers/:id", customerController.Delete)

	orderController := controllers.NewOrderController()
	v1.GET("orders", orderController.GetAll)
	v1.POST("orders", orderController.Create)
	v1.GET("orders/:id", orderController.Detail)
	v1.PUT("orders/:id", orderController.Update)
	v1.DELETE("orders/:id", orderController.Delete)

	return route
}

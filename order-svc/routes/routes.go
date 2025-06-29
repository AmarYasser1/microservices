package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.GetOrder)
	r.GET("/orders", controllers.GetOrders)
	r.PATCH("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)
	r.DELETE("/orders", controllers.DeleteOrders)

    return r;
} 
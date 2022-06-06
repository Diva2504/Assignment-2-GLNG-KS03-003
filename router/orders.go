package router

import (
	"github.com/Diva2504/Assignment-2-GLNG-KS03-003/controllers"
	"github.com/Diva2504/Assignment-2-GLNG-KS03-003/database"
	"github.com/gin-gonic/gin"
)

func OrderRouters() *gin.Engine {
	router := gin.Default()
	db := database.GetDB()
	controller := &controllers.Handler{Connect: db}
	routerGroup := router.Group("/orders")
	{
		routerGroup.GET("/", controller.GetAllOrder)
		routerGroup.POST("/", controller.CreateOrder)
		routerGroup.PUT("/:id", controller.UpdateOrder)
		routerGroup.DELETE("/:id", controller.DeleteOrder)
	}
	return router
}

package routes

import (
	"github.com/Psnsilvino/CaluFestas-Site-api/controllers"
	"github.com/gin-gonic/gin"
)

func ClientRoutes(r *gin.RouterGroup) {
	clients := r.Group("/clients")
	{
		clients.GET("/", controllers.GetClients)
		clients.POST("/", controllers.Register)
	}
}
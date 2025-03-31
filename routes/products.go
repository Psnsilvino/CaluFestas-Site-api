package routes

import (
	"github.com/Psnsilvino/CaluFestas-Site-api/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	clients := r.Group("/products")
	{
		clients.GET("/", controllers.GetProducts)
	}
}
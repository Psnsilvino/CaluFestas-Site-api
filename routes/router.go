package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api") // Agrupa todas as rotas dentro de /api
	{
		ClientRoutes(api)   // Adiciona rotas de usuários
	}

	return router
}
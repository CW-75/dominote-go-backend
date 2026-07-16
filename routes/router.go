package routes

import (
	"github.com/gin-gonic/gin"

	"dominote-go/internal/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	
	// Configurar manejadores
	gameTableHandler := handler.NewGameTableHandler()

	// Rutas de Mesas de Juego
	tables := api.Group("/tables")
	{
		tables.GET("/:uid", gameTableHandler.Get)
		tables.POST("/", gameTableHandler.Create)
		tables.DELETE("/:uid", gameTableHandler.Delete)
	}

	return router
}

func StartRouter(router *gin.Engine) (err error) {
	return router.Run()
}

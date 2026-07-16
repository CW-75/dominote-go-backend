package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GameTableHandler maneja las peticiones HTTP relacionadas con las mesas de juego.
type GameTableHandler struct {
	// TODO: Inyectar casos de uso o repositorios necesarios
}

// NewGameTableHandler crea una nueva instancia de GameTableHandler.
func NewGameTableHandler() *GameTableHandler {
	return &GameTableHandler{}
}

// Get maneja la petición para obtener información de una mesa de juego por su UID.
func (h *GameTableHandler) Get(c *gin.Context) {
	uid := c.Param("uid")
	
	// TODO: Implementar lógica para buscar la mesa en la base de datos usando el uid.

	c.JSON(http.StatusOK, gin.H{
		"message": "Obtener mesa de juego " + uid,
	})
}

// Create maneja la petición para crear una nueva mesa de juego.
func (h *GameTableHandler) Create(c *gin.Context) {
	// TODO: Leer el cuerpo de la petición (JSON) y delegar la creación al caso de uso.
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Mesa de juego creada exitosamente",
	})
}

// Delete maneja la petición para eliminar una mesa de juego por su UID.
func (h *GameTableHandler) Delete(c *gin.Context) {
	uid := c.Param("uid")

	// TODO: Implementar lógica para eliminar la mesa en la base de datos usando el uid.

	c.JSON(http.StatusOK, gin.H{
		"message": "Mesa de juego " + uid + " eliminada",
	})
}

package handlers

import (
	"sistema-oficios/internal/services"

	"github.com/gin-gonic/gin"
)

func ObtenerOficios(c *gin.Context) {
	oficios, err := services.ObtenerOficios()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, oficios)
}

package handlers

import (
	"net/http"

	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/service"
	"github.com/gin-gonic/gin"
)

type PerfilHandler struct {
	servicePerfil service.PerfilService
}

func NewPerfilHandler(service service.PerfilService) *PerfilHandler{
	return &PerfilHandler{service}
}

func(ph *PerfilHandler)Create (c *gin.Context){
 
	var perfilAux models.Perfil
	bodyErr := c.ShouldBindJSON(&perfilAux)
 
	if bodyErr != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Message": bodyErr.Error(),
		})
		return
	}

	err := ph.servicePerfil.Create(&perfilAux)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Message": bodyErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Message": "Usuario creado con exito",
	})


}
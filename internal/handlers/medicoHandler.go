package handlers

import (
	"net/http"

	"github.com/NachooSR/GoToHospital/internal/service"
	"github.com/gin-gonic/gin"
)

type MedicoHandler struct {
	service service.MedicoService
}

func NewMedicoHandler (service service.MedicoService) *MedicoHandler{
	return &MedicoHandler{service}
}

func(mH *MedicoHandler) GetAll(c *gin.Context) {
  
	medicos,err := mH.service.GetAll()

	if err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"mensaje:":" Error al traer los datos",
		})
		return
	}
	c.JSON(200,gin.H{
		"medicos": medicos,
	})

}
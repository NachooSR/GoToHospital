package routes

import (
	"github.com/NachooSR/GoToHospital/internal/handlers"
	"github.com/gin-gonic/gin"
)

func GetAll(rg *gin.RouterGroup,handler *handlers.MedicoHandler){

	rg.GET("/medicos",handler.GetAll)
	
}


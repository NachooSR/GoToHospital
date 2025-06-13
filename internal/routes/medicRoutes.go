package routes

import (
	"github.com/NachooSR/GoToHospital/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RutasMedicos(rg *gin.RouterGroup, handler *handlers.MedicoHandler) {

	//Metodos POST
	rg.POST("medicos/CreateMedico", handler.CreateMedico)

	//Metodos GET
	rg.GET("medicos", handler.GetAll)
	rg.GET("medicos/:id", handler.GetMedicoById)

	rg.DELETE("medicos/:id", handler.DeleteMedico)
	rg.PUT("medicos/:id",handler.Update)

}

package routes

import (
	"github.com/NachooSR/GoToHospital/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RutasMedicos(rg *gin.RouterGroup,handler *handlers.MedicoHandler){

	//Metodos GET
	rg.GET("medicos",handler.GetAll)
	rg.GET("medicos/:id",handler.GetMedicoById) 
	rg.GET("medicos/especialidades",handler.ObtenerMedicosConEspecialidad)
	
	//Metodos POST
	rg.POST("CreateMedico",) //Esto va llamar a un handler que se encargue
}


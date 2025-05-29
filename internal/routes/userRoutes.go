package routes

import (
	"github.com/NachooSR/GoToHospital/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegistrarRutasUsuarios(rg *gin.RouterGroup,handler *handlers.UsuarioHandler){
 
	rg.GET("usuarios",handler.GetAll)
	rg.POST("createUser",handler.CreateUsuario)
	
}
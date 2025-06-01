package routes

import (
	"github.com/NachooSR/GoToHospital/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegistrarRutasUsuarios(rg *gin.RouterGroup,handler *handlers.UsuarioHandler){
 
	rg.POST("createUser",handler.CreateUsuario)
	rg.GET("usuarios",handler.GetAll)
	rg.GET("usuario/:id",handler.GetUserById)
	
}
package routes

import (
	"github.com/NachooSR/GoToHospital/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegistrarRutasUsuarios(rg *gin.RouterGroup,handler *handlers.UsuarioHandler){
 
	rg.POST("usuarios/createUser",handler.CreateUsuario)
	rg.GET("usuarios",handler.GetAll)
	rg.GET("usuario/:id",handler.GetUserById)
	rg.PUT("usuario/:id",handler.Update)
	rg.DELETE("usuario/:id",handler.Delete)
}
package routes

import (
	"github.com/NachooSR/GoToHospital/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterPerfilRoutes(rg *gin.RouterGroup,handler *handlers.PerfilHandler){
   
	rg.POST("perfil/CreatePerfil",handler.Create)

}
package handlers

import (
	"net/http"

	"github.com/NachooSR/GoToHospital/internal/dto"
	"github.com/NachooSR/GoToHospital/internal/models"

	"github.com/NachooSR/GoToHospital/internal/service"
	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	servicio service.UserService
}

func NewHandlerUser (sr service.UserService)*UsuarioHandler {
	return &UsuarioHandler{sr}
}

func (handler *UsuarioHandler)CreateUsuario(c *gin.Context){

	var dtoUser dto.UserDto
	err := c.ShouldBindJSON(&dtoUser)
	
	if err !=nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"Lo sentimos explot",
		})
	}

	usuario := models.Usuario{
		IdRol: dtoUser.IdRol,
		UserName: dtoUser.UserName,
		Password: dtoUser.Password,
	}
	
	id,errorcito:= handler.servicio.CreateUser(&usuario)
	
	if errorcito !=nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"Lo sentimos explot",
		})
	}

	c.JSON(200, gin.H{
		"Id": id,
	})
}

func(userHandler *UsuarioHandler) GetAll(c *gin.Context){

	usuarios,err:= userHandler.servicio.GetAll()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"Lo sentimos explot",
		})
	}
	c.JSON(200,gin.H{
		"Usuarios": usuarios,
	})

}
package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/NachooSR/GoToHospital/internal/dto"

	"github.com/NachooSR/GoToHospital/internal/validations"
	"github.com/NachooSR/GoToHospital/pkg/utils"

	"github.com/NachooSR/GoToHospital/internal/service"
	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	servicio service.UserService
}

func NewHandlerUser(sr service.UserService) *UsuarioHandler {
	return &UsuarioHandler{sr}
}

// /IMPLEMENTACION DE METODOS
func (handler *UsuarioHandler) CreateUsuario(c *gin.Context) {

	var dtoUser dto.UserDto
	err := c.ShouldBindJSON(&dtoUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al parsear los datos",
		})
	}

	//Verificar que los datos no esten vacios
	UsernameCargado := validations.EmptyField(dtoUser.UserName)
	PasswordCargada := validations.EmptyField(dtoUser.Password)

	if UsernameCargado == 1 || PasswordCargada == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Campo contrasena o email vacios",
		})
		return
	}

	//Verificar mail y rigidez de contrasenia
	if !validations.ValidarPassword(dtoUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Pruebe otra contrasenia",
		})
		return
	}
	if !validations.ValidadUsername(dtoUser.UserName) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "El mail es invalido",
		})
		return
	}

	//Si estan los campos llenos y cumplen con los requisitos los derivamos al service
	usuario := dto.DtoToUser(&dtoUser)

	id, errorcito := handler.servicio.CreateUser(&usuario)

	if errorcito != nil {
		if errors.Is(errorcito, utils.ErrUsernameExists) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "El username ya existe"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id,
		"user": usuario})

}

func (userHandler *UsuarioHandler) GetUserById(c *gin.Context) {

	idUrl := c.Param("id")
	idNumber, _ := strconv.Atoi(idUrl)
	usuarioAux, err := userHandler.servicio.GetUserById(idNumber)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "Id no existente",
		})
		return
	}
	usuarioDto := dto.ToUserDtoResponse(usuarioAux)

	c.JSON(http.StatusOK, gin.H{
		"encontrado el user": usuarioDto,
	})
}

func (userHandler *UsuarioHandler) GetAll(c *gin.Context) {

	usuarios, err := userHandler.servicio.GetAll()

	usuariosDTO := dto.ArrayUsersToDto(usuarios)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lo sentimos explot",
		})
	}
	c.JSON(200, gin.H{
		"Usuarios": usuariosDTO,
	})
}

// func(userHandler *UsuarioHandler)Update(c *gin.Context){

// 	idUrl := c.Param("id_user")
//     idNumber,_ := strconv.Atoi(idUrl)

// 	var dtoUser dto.UserDto
// 	errorcito := c.ShouldBindJSON(&dtoUser)

// 	usuarioAux,err := userHandler.servicio.UpdateUser(idNumber,dtoUser)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest,gin.H{
// 			"mensaje":"Your so fool",
// 		})
// 	}
// 	c.JSON(http.StatusOK,gin.H{
// 		"encontrado el user":usuarioAux,
// 	})

// }

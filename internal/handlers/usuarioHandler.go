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

	if UsernameCargado || PasswordCargada {
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

func (userHandler *UsuarioHandler) Update(c *gin.Context) {

	idParam := c.Param("id")
	idUser, _ := strconv.Atoi(idParam)

	//VALIDAR QUE EXISTE USER CON ESE ID
	exist, err := userHandler.servicio.ExistUser(idUser)

	if !exist {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Id no existente",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error de programa ",
		})
		return
	}

	//SI EXISTE
	campos := make(map[string]any)
	errBind := c.ShouldBindJSON(&campos)

	if errBind != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error de programa Binding",
		})
	}

	_, passwordExist := campos["password"]
	_, usernameExist := campos["username"]

	//Validacion Password
	if passwordExist {
		empty := validations.EmptyField(campos["password"].(string))
		if empty {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Password Vacia",
			})
			return
		}
		segura := validations.ValidarPassword(campos["password"].(string))
		if !segura {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Password Debil",
			})
			return
		}
	}

	//Validacion Username
	if usernameExist {
		empty := validations.EmptyField(campos["username"].(string))
		if empty {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Mail Vacio",
			})
			return
		}
		segura := validations.ValidadUsername(campos["username"].(string))
		if !segura {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Formato de mail no valido",
			})
			return
		}
	}

	usuario, errorUpdate := userHandler.servicio.UpdateUser(idUser, campos)

	if errorUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error en el update de datos",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Modificado": usuario,
	})
}

func (userHandler *UsuarioHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	idUser, _ := strconv.Atoi(idParam)

	//VALIDAR QUE EXISTE USER CON ESE ID
	exist, err := userHandler.servicio.ExistUser(idUser)

	if !exist {

		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Id no existente",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error de programa ",
		})
		return
	}

	usuario, _ := userHandler.servicio.GetUserById(idUser)

	err = userHandler.servicio.DeleteRol(usuario.IdUser, usuario.IdRol)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Usuario eliminado (o dado de baja)",
	})

}

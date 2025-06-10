package handlers

import (
	"net/http"
	"strconv"

	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/service"
	"github.com/NachooSR/GoToHospital/internal/validations"
	"github.com/gin-gonic/gin"
)

type MedicoHandler struct {
	service service.MedicoService
}

func NewMedicoHandler(service service.MedicoService) *MedicoHandler {
	return &MedicoHandler{service}
}

// METODOS
func (mh *MedicoHandler) CreateMedico(c *gin.Context) {

	var medicToParse models.Medico
	err := c.ShouldBindJSON(&medicToParse)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Error to Parse data",
		})
		return
	}

	if fieldEmpty := validations.EmptyField(medicToParse.Nombre); fieldEmpty {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Nombre vacio",
		})
		return
	}

	if fieldEmpty := validations.EmptyField(medicToParse.Matricula); fieldEmpty {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Matricula vacia",
		})
		return
	}

	///Verificar si existe la matricula
	matriculaExist, err := mh.service.ExistMatricula(medicToParse.Matricula)

	if matriculaExist {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ERROR": "ERROR EN LA BUSQUEDA DE LA MATRICULA",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": "La matricula ya existe",
		})
		return
	}

	IdMedico, errorCreate := mh.service.Create(&medicToParse)

	if errorCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": errorCreate.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Medico creado",
		"Id":      IdMedico,
	})
}

func (mH *MedicoHandler) GetAll(c *gin.Context) {

	medicosDTO, err := mH.service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"Error": err,
			})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"Medico:": medicosDTO,
			})
	}

}

func (mh *MedicoHandler) GetMedicoById(c *gin.Context) {

	id := c.Param("id")
	numeroID, _ := strconv.Atoi(id)
	medicos, err := mh.service.GetMedicoById(numeroID)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"Error": err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"Medico:": medicos,
		})
}

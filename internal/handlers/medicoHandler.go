package handlers

import (
	"net/http"
	"strconv"

	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/service"
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


	//ID del usuario--> (username y password previamente cargados)

	/*
	  Datos que necesito del medico
	  especialidad //Verificar que existe
	  nombre      //Verificar emptyField
      Matricula  //Empty field Y no repetida
	  Estado    //No verificacion -->Default 
	*/
    var medicToParse models.Medico
	err := c.ShouldBindJSON(&medicToParse)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Message": "Error to Parse data",
		})
		return
	}

	
	

}

func (mH *MedicoHandler) GetAll(c *gin.Context) {

	medicos, err := mH.service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje:": " Error al traer los datos",
		})
		return
	}
	c.JSON(200, gin.H{
		"medicos": medicos,
	})

}

func (mh *MedicoHandler) GetMedicoById(c *gin.Context) {

	id := c.Param("id")
	numeroID, _ := strconv.Atoi(id)
	medicos, err := mh.service.GetMedicoById(numeroID)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"Error": err,
			})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"Medico:": medicos,
			})
	}

}

func (mh *MedicoHandler) ObtenerMedicosConEspecialidad(c *gin.Context) {

	medicosEspecialistas, err := mh.service.ObtenerMedicosConEspecialidad()

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"Error": err,
			})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"Medico:": medicosEspecialistas,
			})
	}
}

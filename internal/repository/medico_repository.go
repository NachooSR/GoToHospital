package repository

import (
	

	"github.com/NachooSR/GoToHospital/internal/models"
)

type MedicoRepository interface {
  
	Crear(medico *models.Medico) error
	GetAll() ([]models.Medico,error)

}
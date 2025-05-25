package repository

import (
	"github.com/NachooSR/GoToHospital/internal/models"
	"gorm.io/gorm"
)

type MedicoRepository interface {

	// Crear(medico *models.Medico) error
	GetAll() ([]models.Medico,error)

}

type medicoRepo struct{
	db *gorm.DB
}

func NewMedicoRepository (db *gorm.DB) MedicoRepository{
	return &medicoRepo{db:db}
}

func (mr *medicoRepo)GetAll()([]models.Medico,error){
   
	var medicos []models.Medico
	err := mr.db.Find(&medicos).Error
	return medicos,err
}
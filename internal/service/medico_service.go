package service

import (
	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
)

type MedicoService interface {
	GetAll() ([]models.Medico,error)
}

type medicoServiceRepo struct{
	repositorio repository.MedicoRepository
}

//Constructor
func NewMedicoService (repo repository.MedicoRepository)MedicoService{
  return &medicoServiceRepo{repo}
}


//sr = serviceRepo
func (sr *medicoServiceRepo) GetAll()([]models.Medico,error){
  return sr.repositorio.GetAll()
}
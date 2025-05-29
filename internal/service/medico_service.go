package service

import (
	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/internal/dto"
)

type MedicoService interface {
	GetAll() ([]models.Medico,error)
	GetMedicoById(int)(models.Medico,error)
	ObtenerMedicosConEspecialidad()([]dto.MedicoDto,error)
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

func (sr *medicoServiceRepo)GetMedicoById(id int)(models.Medico,error){
	return sr.repositorio.GetMedicoById(id)
}

func(sr *medicoServiceRepo)ObtenerMedicosConEspecialidad()([]dto.MedicoDto,error){
	return sr.repositorio.ObtenerMedicosConEspecialidad()
}
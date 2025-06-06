package service

import (
	"github.com/NachooSR/GoToHospital/internal/dto"
	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/pkg/utils"
)

type MedicoService interface {
	Create(*models.Medico)(int,error)
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


func(sr *medicoServiceRepo)Create(medic *models.Medico)(int,error){

	booleano,result:= sr.repositorio.ExistMedic(medic.IdUser)
    if !booleano {
		if result !=nil{
			return 0, result
		}
		return 0, utils.ErrRecordNotFound
	}
	return sr.repositorio.Create(medic)
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
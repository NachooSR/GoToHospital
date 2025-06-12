package service

import (
	"github.com/NachooSR/GoToHospital/internal/dto"
	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/pkg/utils"
)

type MedicoService interface {
	
	Create(*models.Medico) (int, error)
	
	GetAll() ([]dto.MedicoDto, error)
	GetMedicoById(int) (dto.MedicoDto, error)
	
	Update(id int,campos map[string]any)error

	ExistMatricula(string) (bool, error)
	ExistMedic(int)(bool,error)
	DeleteMedico(int)error
}

type medicoServiceRepo struct {
	repositorio     repository.MedicoRepository
	repositorioUser repository.UserRepository
}

// Constructor
func NewMedicoService(repo repository.MedicoRepository, repoMedic repository.UserRepository) MedicoService {
	return &medicoServiceRepo{repositorio: repo, repositorioUser: repoMedic}
}

func (sr *medicoServiceRepo) Create(medic *models.Medico) (int, error) {

	userAux, err := sr.repositorioUser.GetUserById(medic.IdUser)

	if err != nil { //No lo encontro
		return 0, utils.ErrRecordNotFound
	}

	//No es un usuario creado para medico
	if userAux.IdRol != 2 {
		return 0, utils.ErrInvalidRol
	}

	existMedico, errConsulta := sr.repositorio.ExistMedico(medic.IdUser)

	if existMedico {
		return 0, utils.ErrMedicExist
	}
	if errConsulta != nil {
		return 0, errConsulta
	}

	return sr.repositorio.Create(medic)
}

// sr = serviceRepo
func (sr *medicoServiceRepo) GetAll() ([]dto.MedicoDto, error) {
	return sr.repositorio.GetAll()
}

func (sr *medicoServiceRepo) GetMedicoById(id int) (dto.MedicoDto, error) {
	exist,err := sr.repositorio.ExistMedico(id)
	if !exist {
		return dto.MedicoDto{}, err
	}
	return sr.repositorio.GetMedicoById(id)
}

func(sr *medicoServiceRepo)Update(id int,campos map[string]any)error{
    return sr.repositorio.Update(id,campos)
}



func(sr *medicoServiceRepo)DeleteMedico(id int)error{
	
	existMedico,err := sr.repositorio.ExistMedico(id)

	if !existMedico {
		return err
	}

	errEstado:= sr.repositorio.EstadoMedico(id)
	if errEstado != nil {
		return errEstado
	}
	return sr.repositorio.Delete(id)
}


func (sr *medicoServiceRepo) ExistMatricula(matricula string) (bool, error) {
	return sr.repositorio.ExistMatricula(matricula)
}

func (sr *medicoServiceRepo)ExistMedic(id int)(bool,error){
	return sr.repositorio.ExistMedico(id)
}
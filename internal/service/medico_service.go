package service

import (
	"github.com/NachooSR/GoToHospital/internal/dto"
	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/pkg/utils"
)

type MedicoService interface {
	Create(*models.Medico) (int, error)
	GetAll() ([]models.Medico, error)
	GetMedicoById(int) (models.Medico, error)
	ObtenerMedicosConEspecialidad() ([]dto.MedicoDto, error)
	ExistMatricula(string) (bool, error)
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
func (sr *medicoServiceRepo) GetAll() ([]models.Medico, error) {
	return sr.repositorio.GetAll()
}

func (sr *medicoServiceRepo) GetMedicoById(id int) (models.Medico, error) {
	return sr.repositorio.GetMedicoById(id)
}

func (sr *medicoServiceRepo) ObtenerMedicosConEspecialidad() ([]dto.MedicoDto, error) {
	return sr.repositorio.ObtenerMedicosConEspecialidad()
}

func (sr *medicoServiceRepo) ExistMatricula(matricula string) (bool, error) {
	return sr.repositorio.ExistMatricula(matricula)
}

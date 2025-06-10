package repository

import (
	"errors"

	"github.com/NachooSR/GoToHospital/internal/dto"
	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/pkg/utils"
	"gorm.io/gorm"
)

type MedicoRepository interface {
	
	Create(*models.Medico) (int, error)
	
	GetAll() ([]dto.MedicoDto, error)
	GetMedicoById(int) (dto.MedicoDto, error)

	Delete(int) error
	ExistMatricula(string) (bool, error)
	ExistMedico(int) (bool, error)
}

type medicoRepo struct {
	db *gorm.DB
}

// Constructor
func NewMedicoRepository(db *gorm.DB) MedicoRepository {
	return &medicoRepo{db: db}
}

// Implementacion de metodos

func (mr *medicoRepo) Create(medico *models.Medico) (int, error) {
	result := mr.db.Create(&medico).Error
	return medico.IdUser, result
}

func (mr *medicoRepo) GetAll() ([]dto.MedicoDto, error) {

	var resultados []dto.MedicoDto
	err := mr.db.Table("medicos").
		Select("medicos.nombre, medicos.matricula, especialidads.nombre AS especialidad, usuarios.username").
		Joins("left join especialidads on especialidads.id_especialidad = medicos.id_especialidad").
		Joins("left join usuarios on usuarios.id_user = medicos.id_user").
		Where("medicos.estado = ?", "alta").
		Scan(&resultados).Error
	return resultados, err
}

func (mr *medicoRepo) GetMedicoById(id int) (dto.MedicoDto, error) {

	medicoResponse := dto.MedicoDto{}
	err := mr.db.Table("medicos").
	Select("medicos.nombre, medicos.matricula, especialidads.nombre AS especialidad, usuarios.username").
	Joins("left join especialidads on especialidads.id_especialidad = medicos.id_especialidad").
	Joins("left join usuarios on usuarios.id_user = medicos.id_user").
	Where("medicos.id_user = ?", id).
	Scan(&medicoResponse).Error

	return medicoResponse, err
}


func (mr *medicoRepo) Delete(id int) error {
	return mr.db.Model(&models.Medico{}).Where("id_user= ?", id).Update("estado", "baja").Error
}

/*
True existe
False y Nil no existe y no hay error
*/
func (mr *medicoRepo) ExistMatricula(matricula string) (bool, error) {

	var medicoAux models.Medico
	result := mr.db.Where("matricula = ?", matricula).First(&medicoAux).Error

	if result == nil { //existe
		return true, nil
	}

	if errors.Is(result, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, result
}

// True existe False y Nil no existe
func (mr *medicoRepo) ExistMedico(id int) (bool, error) {
	var medicoAux models.Medico
	result := mr.db.First(&medicoAux, id).Error

	if result == nil {
		return true, nil
	}

	if errors.Is(result, gorm.ErrRecordNotFound) {
		return false, utils.ErrRecordNotFound
	}

	return false, result
}

package repository

import (
	"github.com/NachooSR/GoToHospital/internal/dto"
	"github.com/NachooSR/GoToHospital/internal/models"
	"gorm.io/gorm"
)

type MedicoRepository interface {
	Create(medico *models.Medico) error
	GetAll() ([]models.Medico, error)
	GetMedicoById(int) (models.Medico, error)
	ObtenerMedicosConEspecialidad() ([]dto.MedicoDto, error)

	Delete(int) error
}

type medicoRepo struct {
	db *gorm.DB
}

// Constructor
func NewMedicoRepository(db *gorm.DB) MedicoRepository {
	return &medicoRepo{db: db}
}

// Implementacion de metodos
func (mr *medicoRepo) Create(medico *models.Medico) error {
	var errorx error
	return errorx
}

func (mr *medicoRepo) GetAll() ([]models.Medico, error) {

	var medicos []models.Medico
	err := mr.db.Find(&medicos).Error
	return medicos, err
}

func (mr *medicoRepo) GetMedicoById(id int) (models.Medico, error) {

	var medicoaux = models.Medico{
		IdUser: id,
	}
	result := mr.db.First(&medicoaux, id)

	return medicoaux, result.Error
}

func (mr *medicoRepo) ObtenerMedicosConEspecialidad() ([]dto.MedicoDto, error) {
	var resultados []dto.MedicoDto
	err := mr.db.Table("medicos").
		Select("medicos.nombre, medicos.matricula, especialidads.nombre AS especialidad").
		Joins("left join especialidads on especialidads.id_especialidad = medicos.id_especialidad").
		Scan(&resultados).Error
	return resultados, err
}

func (mr *medicoRepo) Delete(id int) error {
	return mr.db.Model(&models.Medico{}).Where("id_user= ?", id).Update("estado", "baja").Error
}

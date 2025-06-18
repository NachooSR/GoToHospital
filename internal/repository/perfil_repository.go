package repository

import (
	"errors"

	"github.com/NachooSR/GoToHospital/internal/models"
	"gorm.io/gorm"
)

type PerfilRepository interface {
	CreatePerfil(*models.Perfil) error
	ExistDNI(string)(bool,error)
}

type perfilRepo struct {
	db *gorm.DB
}

func NewPerfilRepository(db *gorm.DB) PerfilRepository {
	return &perfilRepo{db}
}


func(pr *perfilRepo) CreatePerfil(auxPerfil *models.Perfil) error {

	result := pr.db.Create(&auxPerfil).Error
	return result
}

func(pr *perfilRepo)ExistDNI(dni string)(bool,error){
	var perfilAux models.Perfil
	result := pr.db.Where("documento = ?", dni).First(&perfilAux).Error

	if result == nil { //existe
		return true, nil
	}

	if errors.Is(result, gorm.ErrRecordNotFound) {
		return false, gorm.ErrRecordNotFound
	}

	return false, result
}

func (pr *perfilRepo)ExistNroSocio(nroSocio string)(bool,error){
	var perfilAux models.Perfil
	result := pr.db.Where("nro_carnet = ?", nroSocio).First(&perfilAux).Error

	if result == nil { //existe
		return true, nil
	}

	if errors.Is(result, gorm.ErrRecordNotFound) {
		return false, gorm.ErrRecordNotFound
	}

	return false, result
}
package validations

import (
	"errors"

	"github.com/NachooSR/GoToHospital/internal/models"
	"gorm.io/gorm"
)

// Esta funcion -->servicio o repo
func ValidateIdRol(id int, conex *gorm.DB) (bool, error) {

	RolAuxiliar := models.Rols{}
	result := conex.First(&RolAuxiliar, id).Error

	if result == nil {
		return true, nil
	}

	if errors.Is(result, gorm.ErrRecordNotFound) {
		//Error = nil porque no hay error, sino que no existe
		return false, nil
	}

	return false, result

}




package repository

import (
	"github.com/NachooSR/GoToHospital/internal/models"
	"gorm.io/gorm"
)

//La interfaz es la que guarda los metodos
type UserRepository interface {
	CreateUser(*models.Usuario) (int, error)
	GetAll()([]models.Usuario,error)
}

type userRepoGorm struct {
	db *gorm.DB //Conex a la db
}

func NewUserRepo (conex  *gorm.DB) UserRepository {
	return &userRepoGorm{conex}
}

func (r *userRepoGorm) CreateUser(aux *models.Usuario) (int, error) {

	result := r.db.Create(&aux)

	//Variables de retorno
	ID_user:= aux.IdUser
	resultError:= result.Error 

	return ID_user,resultError
}

func (r *userRepoGorm)GetAll()([]models.Usuario,error){

	var users []models.Usuario
	err := r.db.Find(&users).Error
	return users,err
}
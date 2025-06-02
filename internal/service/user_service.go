package service

import (

	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/pkg/utils"
)

type UserService interface {
	CreateUser(*models.Usuario) (int, error)
	GetAll() ([]models.Usuario, error)
	GetUserById(int) (*models.Usuario, error)
	UpdateUser(int, *models.Usuario) (*models.Usuario, error)
}

type userServiceRepo struct {
	repo repository.UserRepository
}



func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceRepo{repo}
}

func (sr *userServiceRepo) CreateUser(auxiliar *models.Usuario) (int, error) {

	
	exists, err := sr.repo.ExistUsername(auxiliar.UserName)
    if err != nil {
        return 0, err // Error 
    }
    if exists {
        return 0, utils.ErrUsernameExists // No hay error, pero el usuario ya existe
    }

    return sr.repo.CreateUser(auxiliar) // Crear el usuario

}

func (sr *userServiceRepo) GetUserById(id int) (*models.Usuario, error) {
	return sr.repo.GetUserById(id)
}

func (sr *userServiceRepo) GetAll() ([]models.Usuario, error) {
	return sr.repo.GetAll()
}

func (sr *userServiceRepo) UpdateUser(id int, user *models.Usuario) (*models.Usuario, error) {
	return sr.repo.Update(id, user)
}

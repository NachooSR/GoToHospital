package service

import (
	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
)

type UserService interface {
	CreateUser(*models.Usuario)(int,error)
	GetAll()([]models.Usuario,error)
}

type userServiceRepo struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository)UserService{
	return &userServiceRepo{repo}
}

func (sr *userServiceRepo)CreateUser(auxiliar *models.Usuario)(int,error){

	//Aca hariamos la validacion--> username (email) password (seguridad)
	
	//Derivacion al Repositorio
    return sr.repo.CreateUser(auxiliar)
}

func (sr *userServiceRepo)GetAll()([]models.Usuario,error){
	return sr.repo.GetAll()
}
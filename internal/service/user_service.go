package service

import (
	"errors"

	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/pkg/utils"
)

type UserService interface {
	CreateUser(*models.Usuario) (int, error)
	
	GetUserById(int) (*models.Usuario, error)
	
	GetAll() ([]models.Usuario, error)

	UpdateUser(int,map[string]any) (*models.Usuario, error)

    DeleteUser(int)error

	DeleteMedico(int)error

	DeleteRol(int,int)error

	ExistUser(int)(bool,error)
}

type userServiceRepo struct {
	repo repository.UserRepository
	medicoRepo repository.MedicoRepository
}

func NewUserService(repo repository.UserRepository, medicoRep repository.MedicoRepository) UserService {
	return &userServiceRepo{repo:repo,
		medicoRepo:medicoRep}
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

func (sr *userServiceRepo) UpdateUser(id int, campos map[string]any) (*models.Usuario, error) {
	return sr.repo.Update(id,campos)
}

func(sr *userServiceRepo)DeleteUser(id int)error{
	return sr.repo.Delete(id)
}

func(sr *userServiceRepo)DeleteMedico(id int)error{
	return sr.medicoRepo.Delete(id)
}

func(sr *userServiceRepo)DeleteRol(id,rol int)error{

	switch rol{
	case 2: return sr.medicoRepo.Delete(id)
	case 3: return nil
	case 4: return nil
	default: return errors.New("PRUEBA")
	}
}


func(sr *userServiceRepo) ExistUser(id int)(bool,error){
 return sr.repo.ExistIdUser(id)
}
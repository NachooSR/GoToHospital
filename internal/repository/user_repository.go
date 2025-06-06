package repository

import (
	"errors"

	"github.com/NachooSR/GoToHospital/internal/models"

	"gorm.io/gorm"
)

// La interfaz es la que guarda los metodos
type UserRepository interface {

	//Create-->Get (all/id)-->Update-->Delete

	CreateUser(*models.Usuario) (int, error)

	GetUserById(int) (*models.Usuario, error)
	GetAll() ([]models.Usuario, error)

	Update(int, map[string]any) (*models.Usuario, error)

	Delete(int) error

	ExistUsername(string) (bool, error)
	ExistIdUser(int) (bool, error)
}

type userRepoGorm struct {
	db *gorm.DB
}

func NewUserRepo(conex *gorm.DB) UserRepository {
	return &userRepoGorm{conex}
}

///IMPLEMENTACION DE METODOS

func (r *userRepoGorm) CreateUser(aux *models.Usuario) (int, error) {

	result := r.db.Create(aux)
	return aux.IdUser, result.Error
}

func (r *userRepoGorm) GetUserById(id int) (*models.Usuario, error) {

	var aux models.Usuario
	result := r.db.First(&aux, id).Error

	if result != nil {
		return nil, result
	}

	return &aux, nil
}

func (r *userRepoGorm) GetAll() ([]models.Usuario, error) {

	var users []models.Usuario
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepoGorm) Update(id int, campos map[string]any) (*models.Usuario, error) {
	var usuario models.Usuario

	//Esto es para cargar el usuario y poder devolverlo, si llego aca es porque paso la verificacion del servicio donde existia el ID
	err := r.db.First(&usuario, id).Error
	if err != nil {
		return nil, err
	}

	err2 := r.db.Model(&usuario).Updates(campos).Error
	if err2 != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *userRepoGorm) Delete(id int) error {

	user := models.Usuario{}
	return r.db.Delete(&user, id).Error
}

// /FUNCIONES AUXILIARES
func (r *userRepoGorm) ExistUsername(username string) (bool, error) {

	usuarioAux := models.Usuario{}
	result := r.db.Where("username = ?", username).First(&usuarioAux).Error

	if result == nil {
		return true, nil
	}

	if errors.Is(result, gorm.ErrRecordNotFound) {
		//Error = nil porque no hay error, sino que no existe
		return false, nil
	}

	return false, result
}

func (r *userRepoGorm) ExistIdUser(id int) (bool, error) {
	usuarioAux := models.Usuario{}
	result := r.db.First(&usuarioAux).Error

	if result == nil {
		return true, nil
	}

	if errors.Is(result, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, result
}

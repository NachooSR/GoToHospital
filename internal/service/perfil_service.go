package service

import (
	"errors"

	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/pkg/validations"
)

type PerfilService interface {
	Create(*models.Perfil) error
}

type perfilSer struct {
	repo repository.PerfilRepository
}

func NewServicePerfil(repo repository.PerfilRepository) PerfilService {
	return &perfilSer{repo: repo}
}

func (ps *perfilSer) Create(auxPerfil *models.Perfil) error {

	// ID, NOMBRE, APELLIDO, DIRECCION, NROTELEFONO, DNI, NROSOCIO
	//Existe el id

	//Campo vacio
	if nombre := validations.EmptyField(auxPerfil.Nombre); nombre {
		return errors.New("nombre Vacio")
	}
	if apellido := validations.EmptyField(auxPerfil.Nombre); apellido {
		return errors.New("apellido Vacio")
	}
	if direccion := validations.EmptyField(auxPerfil.Direccion); direccion {
		return errors.New("direccion Vacia")
	}
	if nroTelefono := validations.EmptyField(auxPerfil.NroTelefono); nroTelefono {
		return errors.New("nro Telefono Vacio")
	}
	if dni := validations.EmptyField(auxPerfil.Documento); dni {
		return errors.New("dni Vacio")
	}
	if nroSocio := validations.EmptyField(auxPerfil.NroCarnet); nroSocio {
		return errors.New("numero carnet Vacio")
	}

	//Validar que no exista DNI
	exist,err := ps.repo.ExistDNI(auxPerfil.Documento)
	if exist {
		return err
	}

	//Validar que no exista nro_socio
	exist2,err2 := ps.repo.ExistDNI(auxPerfil.Documento)
	if exist2 {
		return err2
	}

	return ps.repo.CreatePerfil(auxPerfil)
}

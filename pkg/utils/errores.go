package utils

import "errors"


//Username errors
var ErrUsernameExists = errors.New("username ya existe")

var ErrRecordNotFound = errors.New("ID username no existe")

var ErrInvalidRol = errors.New("usuario no tiene rol correspondiente")


//Medico errors

var ErrMedicExist = errors.New("medico ya existente")

var ErrBaja = errors.New("el medico ya se encuentra dado de baja")
package utils

import "errors"

var ErrUsernameExists = errors.New("username ya existe")

var ErrRecordNotFound = errors.New("ID username no existe")

var ErrInvalidRol = errors.New("El usuario no tiene rol correspondiente")

var ErrMedicExist = errors.New("El medico ya existe")

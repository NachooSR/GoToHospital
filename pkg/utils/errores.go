package utils

import "errors"

var ErrUsernameExists = errors.New("username ya existe")

var ErrRecordNotFound = errors.New("ID username no existe")
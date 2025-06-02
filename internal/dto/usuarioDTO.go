package dto

import "github.com/NachooSR/GoToHospital/internal/models"

type UserDto struct {
	IdRol    int    `json:"id_rol"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserDtoResponse struct {
	IdUser   int    `json:"id_user"`
	IdRol    int    `json:"id_rol"`
	UserName string `json:"username"`
}

func DtoToUser(aux *UserDto) models.Usuario {
	return models.Usuario{
		IdRol:    aux.IdRol,
		UserName: aux.UserName,
		Password: aux.Password,
	}
}

func ToUserDtoResponse(aux *models.Usuario) UserDtoResponse {
	return UserDtoResponse{
		IdUser:   aux.IdUser,
		IdRol:    aux.IdRol,
		UserName: aux.UserName,
	}
}

func ArrayUsersToDto(usuarios []models.Usuario) []UserDtoResponse {
	usuariosDTO := make([]UserDtoResponse, len(usuarios))
	for i, usuario := range usuarios {
		usuariosDTO[i] =ToUserDtoResponse(&usuario)
	}
	return usuariosDTO
}

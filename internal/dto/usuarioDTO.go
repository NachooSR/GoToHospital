package dto

type UserDto struct {
	IdRol    int    `json:"id_rol"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserDtoResponse struct{
    IdUser   int    `json:"id_user"`
	IdRol    int    `json:"id_rol"`
	UserName string `json:"username"`
}
package dto

type UserDto struct {
	IdRol    int    `json:"id_rol"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

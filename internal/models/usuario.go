package models

type Usuario struct {
	IdUser   int    `gorm:"column:id_user;primaryKey;autoIncrement" json:"id_user"`
	IdRol    int    `gorm:"column:id_rol" json:"id_rol"`
	UserName string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
type Usuarios[]Usuario
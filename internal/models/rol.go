package models

type Rol struct {
	IdRol  int    `gorm:"column:id_rol" json:"id_rol"`
	Numero int    `gorm:"column:numero" json:"numero"`
	Nombre string `gorm:"column:nombre" json:"nombre"`
}
type Rols[]Rol
package models

type Perfil struct {
	IDUser      int    `gorm:"column:id_user;primaryKey" json:"id_user"`
	Nombre      string `gorm:"column:nombre" json:"nombre"`
	Apellido    string `gorm:"column:apellido" json:"apellido"`
	NroTelefono string `gorm:"column:nro_telefono" json:"nro_telefono"`
	Direccion   string `gorm:"column:direccion" json:"direccion"`
	Documento   string `gorm:"column:documento" json:"documento"`
	NroCarnet   string `gorm:"column:nro_carnet" json:"nro_carnet"`
	Estado      string `gorm:"column:estado" json:"estado"` // enum: estado_perfil
}
type Perfils[]Perfil
package models

type Consultorio struct {
	IDConsultorio int  `gorm:"column:id_consultorio" json:"id_consultorio"`
	Numero        int  `gorm:"column:numero" json:"numero"`
	Piso          int  `gorm:"column:piso" json:"piso"`
	Ocupado       bool `gorm:"column:ocupado" json:"ocupado"`
}
type Consultorios[]Consultorio
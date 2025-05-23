package models

type Especialidad struct {
	IdEspecialidad int    `gorm:"column:id_especialidad" json:"id_especialidad"`
	Nombre         string `gorm:"column:nombre" json:"nombre"`
	Numero         int    `gorm:"column:numero" json:"numero"`
}

type Especialidads[]Especialidad
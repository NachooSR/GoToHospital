package models

type Medico struct {
	IdUser         int    `gorm:"column:id_user" json:"id_user"`
	IdEspecialidad int    `gorm:"column:id_especialidad" json:"id_especialidad"`
	Nombre         string `gorm:"column:nombre" json:"nombre"`
	Matricula      string `gorm:"column:matricula" json:"matricula"`

	Especialidad Especialidad `gorm:"foreignKey:IdEspecialidad;references:IdEspecialidad"`
}
type Medicos[]Medico
package models

type Medico struct {
	IdUser         int          `gorm:"column:id_user;primaryKey" json:"id_user"`
	IdEspecialidad int          `gorm:"column:id_especialidad" json:"id_especialidad"`
	Nombre         string       `gorm:"column:nombre" json:"nombre"`
	Matricula      string       `gorm:"column:matricula" json:"matricula"`
	Estado         string       `gorm:"column:estado" json:"estado"`
	Especialidad   Especialidad `gorm:"foreignKey:IdEspecialidad;references:IdEspecialidad"`
}


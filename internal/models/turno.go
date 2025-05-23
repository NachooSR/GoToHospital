package models

import "time"

type Turno struct {
	IDTurno       int       `gorm:"column:id_turno" json:"id_turno"`
	Dia           time.Time `gorm:"column:dia" json:"dia"` 
	Hora          time.Time `gorm:"column:hora" json:"hora"` 
	IDMedico      int       `gorm:"column:id_medico" json:"id_medico"`
	IDConsultorio int       `gorm:"column:id_consultorio" json:"id_consultorio"`
	IDUser        int       `gorm:"column:id_user" json:"id_user"`
	Estado        string    `gorm:"column:estado" json:"estado"` 
}
type Turnos[]Turnos
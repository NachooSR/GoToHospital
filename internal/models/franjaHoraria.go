package models

import "time"

type FranjaHoraria struct {
	IDFranja      int       `gorm:"column:id_franja" json:"id_franja"`
	IDUser        int       `gorm:"column:id_user" json:"id_user"`
	Dia           string    `gorm:"column:dia" json:"dia"`
	HoraInicio    time.Time `gorm:"column:hora_inicio" json:"hora_inicio"`
	HoraCierre    time.Time `gorm:"column:hora_cierre" json:"hora_cierre"`
	IDConsultorio int       `gorm:"column:id_consultorio" json:"id_consultorio"`
}
type FranjaHorarias[]FranjaHoraria
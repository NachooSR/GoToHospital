package models

type Visita struct {
	IdVisita    int    `gorm:"column:id_visita" json:"id_visita"`
	IdTurno     int    `gorm:"column:id_turno" json:"id_turno"`
	Asunto      string `gorm:"column:asunto" json:"asunto"`
	Descripcion string `gorm:"column:descripcion" json:"descripcion"`
}

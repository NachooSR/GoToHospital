package models
import "time"

type SolicitudRegistro struct {
	IDSolicitud  int       `gorm:"column:id_solicitud" json:"id_solicitud"`
	IDUser       int       `gorm:"column:id_user" json:"id_user"`
	Estado       string    `gorm:"column:estado" json:"estado"`
	FechaEmision time.Time `gorm:"column:fecha_emision" json:"fecha_emision"`
}
type SolicitudRegistros[]SolicitudRegistro
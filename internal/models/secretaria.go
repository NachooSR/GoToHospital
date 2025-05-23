package models

type Secretaria struct {
	IdUser int `gorm:"column:id_user" json:"id_user"`
}
type Secretarias[]Secretaria
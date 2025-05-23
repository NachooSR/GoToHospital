package models

type Admin struct {
	IdUser int `gorm:"column:id_user" json:"id_user"`
}
type Admins[]Admin
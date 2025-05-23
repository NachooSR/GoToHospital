package models

type HistoriaClinica struct {
	IDUser                 int     `gorm:"column:id_user" json:"id_user"`
	Peso                   float64 `gorm:"column:peso" json:"peso"`
	Altura                 float64 `gorm:"column:altura" json:"altura"`
	Edad                   int     `gorm:"column:edad" json:"edad"`
	Alergias               bool    `gorm:"column:alergias" json:"alergias"`
	AlergiasDetalle        string  `gorm:"column:alergias_detalle" json:"alergias_detalle"`
	AntecedentesFamiliares bool    `gorm:"column:antecedentes_familiares" json:"antecedentes_familiares"`
	AntecedentesDetalle    string  `gorm:"column:antecedentes_detalle" json:"antecedentes_detalle"`
	EnfermedadesCardiacas  bool    `gorm:"column:enfermedades_cardiacas" json:"enfermedades_cardiacas"`
	EnfermedadesDetalle    string  `gorm:"column:enfermedades_detalle" json:"enfermedades_detalle"`
	Fuma                   bool    `gorm:"column:fuma" json:"fuma"`
	Medicacion             bool    `gorm:"column:medicacion" json:"medicacion"`
	MedicacionDetalle      string  `gorm:"column:medicacion_detalle" json:"medicacion_detalle"`
	Diabetes               bool    `gorm:"column:diabetes" json:"diabetes"`
	Ejercicio              bool    `gorm:"column:ejercicio" json:"ejercicio"`
}
type HistoriaClinicas[]HistoriaClinica

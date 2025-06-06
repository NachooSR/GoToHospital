package dto

type MedicoDto struct {
	Mail         string `json:"username"`
	Nombre       string `json:"nombre"`
	Especialidad string `json:"especialidad"`
	Matricula    string `json:"matricula"`
}

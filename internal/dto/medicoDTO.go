package dto

type MedicoDto struct {
	Username     string `json:"username"`
	Nombre       string `json:"nombre"`
	Especialidad string `json:"especialidad"`
	Matricula    string `json:"matricula"`
}

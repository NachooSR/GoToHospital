package main

import (
	"fmt"

	"github.com/NachooSR/GoToHospital/internal/config"
	"github.com/NachooSR/GoToHospital/internal/repository"
	// "github.com/NachooSR/GoToHospital/internal/models"
	// "github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/internal/models"
	"github.com/NachooSR/GoToHospital/internal/validations"
	//"gorm.io/gorm"
	//"github.com/NachooSR/GoToHospital/internal/service"
)

func main() {
   
	//Inicializacion de db y env variables
	configuration := config.LoadConfig()
	db := config.ConnectDb(configuration)
	//

	/* Pruebas de medico
		//Repo-->Service-->Handler-->Ruta
		repositorioPruebaMedico := repository.NewMedicoRepository(db)
		servicioPruebaMedico := service.NewMedicoService(repositorioPruebaMedico)
	    //

		var medicos []models.Medico
		var err error
		medicos,err = servicioPruebaMedico.GetAll()

		if err ==nil{
	    for _,medico := range medicos {
			fmt.Println(medico)
		}
	  	}else{
			fmt.Println(err)
		}

		id:=4

		medicoAux,err := servicioPruebaMedico.GetMedicoById(id)
		if err !=nil{
			fmt.Println("Lo sentimos no existe:")
			fmt.Println(err)
		}else{
			fmt.Println(medicoAux)
		}

		medicos2,errorcito:= repositorioPruebaMedico.ObtenerMedicosConEspecialidad()
		if errorcito != nil{
			return
		}
		fmt.Println(medicos2) */



	repoUser := repository.NewUserRepo(db)

	//Cargamos un usuario para probar
	usuarioAux := models.Usuario{
		IdUser: 16,
		IdRol:    3,
		UserName: "holis@gmail.com",
		Password: "Sprinfield.123#",
	}


	
	numerito := validations.EmptyField(usuarioAux.UserName)
	numeroPassword := validations.EmptyField(usuarioAux.Password)

	if numerito != 0 || numeroPassword != 0 {

		fmt.Println("Error campo email o password vacios")
	    return
	}


	if !validations.ValidadUsername(usuarioAux.UserName) {
		fmt.Println("El formato de mail es incorrecto")
		return
	}
	if !validations.ValidarPassword(usuarioAux.Password) {
		fmt.Println("Error la contrasena no es segura")
		return
	}

	fmt.Println("Fields llenos y seguros :)")



	id, err := repoUser.CreateUser(&usuarioAux)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)

}

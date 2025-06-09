package main

import (

	//"github.com/NachooSR/GoToHospital/internal/config"
	"github.com/NachooSR/GoToHospital/internal/config"
	"github.com/NachooSR/GoToHospital/internal/handlers"
	"github.com/NachooSR/GoToHospital/internal/repository"
	"github.com/NachooSR/GoToHospital/internal/routes"
	"github.com/NachooSR/GoToHospital/internal/service"
	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

//Aqui deberiamos hacer un init de las variables

//

var prefijo string = "/api/v1/"

func main() {

	// Conexion a la db y carga de .env
	configuration := config.LoadConfig()
	db := config.ConnectDb(configuration)

	///Inicializar Repos
	repositorioMedico:= repository.NewMedicoRepository(db)
	repositorioUsuario := repository.NewUserRepo(db)

	//Inicializar servicios
	servicioUser := service.NewUserService(repositorioUsuario,repositorioMedico)
	servicioMedico:= service.NewMedicoService(repositorioMedico,repositorioUsuario)

    //Inicializar handlers
	handlerUser := handlers.NewHandlerUser(servicioUser)
	handlerMedico:= handlers.NewMedicoHandler(servicioMedico)


	// Inicializar router
	router := gin.Default()
	api := router.Group(prefijo)

	//Registro rutas

	routes.RutasMedicos(api, handlerMedico)
	routes.RegistrarRutasUsuarios(api, handlerUser)

	// Levantar servidor
	router.Run() // default en puerto 8080
}

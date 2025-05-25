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

	// Inicializar repo, service y handler
	repo := repository.NewMedicoRepository(db)
	svc := service.NewMedicoService(repo)
	handler := handlers.NewMedicoHandler(svc)

	// Inicializar router y registrar rutas
	router := gin.Default()
	api := router.Group("/api/v1")
	routes.GetAll(api, handler)

	// Levantar servidor
	router.Run() // default en puerto 8080
}

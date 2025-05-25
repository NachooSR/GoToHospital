package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBServer   string
	DBPort     string
	DBName     string
}

func LoadConfig() *Config {
	err := godotenv.Load("./internal/config/.env")
	if err != nil {
		log.Fatal("Error loading .env ")
	}

	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBServer:   os.Getenv("DB_SERVER"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func ConnectDb (configuracion *Config) *gorm.DB {

	// dsn :=
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println("ERRORCITO LINEA 33 CONEX")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		configuracion.DBServer, configuracion.DBUser, configuracion.DBPassword,
		configuracion.DBName, configuracion.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al abrir la DB")
	}
	return db

}

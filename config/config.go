package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		DBHost:     os.Getenv("DBHOST"),
		DBPort:     os.Getenv("DBPORT"),
		DBUser:     os.Getenv("DBUSER"),
		DBPassword: os.Getenv("DBPASSWORD"),
		DBName:     os.Getenv("DBNAME")}

	return config

}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func GetPort() string {
	var port string = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return ":" + port
}

func GetEnv(key string, defaultValue string) string {
	var value string = os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

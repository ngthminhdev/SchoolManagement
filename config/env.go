package config

import (
	"GolangBackend/internal/global"
	"log"
	"os"
	"strings"

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

func SetWhiteListPaths() {
	var value string = GetEnv("WHILE_LIST_PATHS", "")
	var result map[string]string = map[string]string{}
	if value == "" {
		global.WhileListPaths = result
		return
	}

	for _, item := range strings.Split(value, ";") {
		whileListEnv := strings.Split(item, "_")
		method := whileListEnv[0]
		path := whileListEnv[1]

		result[path] = method
	}

	global.WhileListPaths = result
}

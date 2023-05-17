package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("* Loading Environment Configuration")
	loadEnv()
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + "owltier.com" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Println("* Error loading .env file")
	}
}

func Environment() string {
	return os.Getenv("ENVIRONMENT")
}

func ListenAddr() string {
	return os.Getenv("LISTEN_ADDR")
}

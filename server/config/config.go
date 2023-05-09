package config

import (
	"log"
	"os"
	"regexp"
	"strings"

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

	err := godotenv.Load(string(rootPath) + `/.env.local`)

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

func AccessSecret() []byte {
	return []byte(os.Getenv("JWT_ACCESS_SECRET"))
}

func RefreshSecret() []byte {
	return []byte(os.Getenv("JWT_REFRESH_SECRET"))
}

func EmailTokenSecret() []byte {
	return []byte(os.Getenv("JWT_EMAIL_SECRET"))
}

func JwtIssuer() string {
	return os.Getenv("JWT_ISSUER")
}

func JwtAudience() []string {
	aud := os.Getenv("JWT_AUDIENCE")
	audience := strings.Split(aud, ",")
	return audience
}

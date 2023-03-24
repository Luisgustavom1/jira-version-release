package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	godotenv.Load("./.env")
}

func SetEnv(envName string, value string) {
	os.Setenv(envName, value)
}

func GetEnv(envName string) string {
	return os.Getenv(envName)
}

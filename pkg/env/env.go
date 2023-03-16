package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

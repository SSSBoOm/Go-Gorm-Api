package config

import (
	"os"

	"github.com/joho/godotenv"

	"fmt"
)

func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}

	return os.Getenv(key)

}

package src

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetKey(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error lpading .env %v", err)
	}
	key := os.Getenv(name)
	return key
}

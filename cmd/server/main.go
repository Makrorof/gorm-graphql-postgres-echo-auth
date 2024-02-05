package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/makrorof/gorm-graphql-postgres-echo-auth/internal/app/server"
)

func main() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	server.Run()
}

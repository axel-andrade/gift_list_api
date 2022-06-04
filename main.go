package main

import (
	database "go_gift_list_api/src/infra/database"
	"go_gift_list_api/src/infra/http/server"
	"log"

	"github.com/joho/godotenv"
)

/**
A função init por padrão é a primeira a ser executada pelo go.
Utilizada para configurar ou fazer um pré carregamento.
**/
func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.ConnectDB()
}

func main() {
	server := server.NewServer()
	server.Run()
}

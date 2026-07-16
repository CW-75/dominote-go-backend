package main

import (
	"log"

	"dominote-go/internal/database"
	"dominote-go/routes"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("No se pudo inicializar la base de datos: %v", err)
	}

	router := routes.InitRouter()
	
	log.Println("Iniciando el servidor API...")
	if err := routes.StartRouter(router); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

package main

import (
	"log"

	"api-rest-with-go/internal/config"
	"api-rest-with-go/internal/core/domain"
	"api-rest-with-go/internal/core/ports"
	"api-rest-with-go/internal/core/services"
	"api-rest-with-go/internal/infrastructure/database/postgres"
	"api-rest-with-go/internal/infrastructure/server"
)

func main() {
	// Cargar configuración
	cfg := config.GetConfig()

	// Inicializar base de datos
	db, err := postgres.NewConnection(cfg)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(&domain.Item{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	// Inicializar repositorios
	itemRepo := postgres.NewItemRepository(db)

	// Inicializar servicios
	itemService := services.NewItemService(itemRepo)

	// Inicializar y arrancar servidor
	srv := server.NewServer(cfg, &ports.Service{
		Items: itemService,
	}, db)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}

package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-rest-with-go/internal/config"
	"api-rest-with-go/internal/core/ports"
	"api-rest-with-go/internal/infrastructure/server/http/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	config     *config.Config
	router     *gin.Engine
	httpServer *http.Server
}

func NewServer(cfg *config.Config, services *ports.Service, db *gorm.DB) *Server {
	server := &Server{
		config: cfg,
		router: gin.New(),
	}

	// Configurar rutas
	routes.Setup(server.router, services, db)

	server.httpServer = &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      server.router,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}

	return server
}

func (s *Server) Start() error {
	// Canal para errores del servidor
	errChan := make(chan error, 1)

	// Iniciar servidor en una goroutine
	go func() {
		log.Printf("Server starting on port %s", s.config.Server.Port)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("error starting server: %w", err)
		}
	}()

	// Canal para señales del sistema
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Esperar por error o señal de término
	select {
	case err := <-errChan:
		return err
	case <-quit:
		log.Println("Shutting down server...")
		return s.Shutdown()
	}
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}

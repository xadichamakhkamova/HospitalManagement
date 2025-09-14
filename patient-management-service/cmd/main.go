package main

import (
	app "patient-service/internal/app"
	config "patient-service/internal/config"
	"patient-service/internal/repository"
	pq "patient-service/internal/repository/postgres"
	service "patient-service/internal/service"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"patient-service/logger"
)

func main() {
	// Init logger
	log := logger.NewLogger()

	// Load config
	cfg, err := config.Load("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	log.Info("Configuration loaded successfully")

	// Connect DB
	db, err := pq.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Info("Connected to Postgres database")

	// Repository init
	queries := repository.NewPatientSqlc(db)
	repo := repository.NewIPatientRepository(queries)
	log.Info("Patient repository initialized")

	// Service init
	srv := service.NewPatientService(repo)
	log.Info("Patient service initialized")

	// App init
	application := app.New(*srv, cfg.Service.Port)
	log.Infof("Patient Service is starting on port: %d", cfg.Service.Port)

	// Run server in goroutine
	go func() {
		application.MustRun()
	}()
	log.Info("Server is running...")

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop
	log.Info("Received signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	application.Stop()
	log.Info("Server stopped gracefully")

	<-ctx.Done()
	log.Info("Shutdown complete")
}

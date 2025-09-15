package main

import (
	app "pharmacist-service/internal/app"
	config "pharmacist-service/internal/config"
	"pharmacist-service/internal/repository"
	pq "pharmacist-service/internal/repository/postgres"
	service "pharmacist-service/internal/service"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pharmacist-service/logger" 
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
	queries := repository.NewPharmaSqlc(db, log)
	repo := repository.NewIPharmacistRepository(queries, log)
	log.Info("Pharmacist repository initialized")

	// Service init
	srv := service.NewPharmaService(repo)
	log.Info("Pharmacist service initialized")

	// App init
	application := app.New(*srv, cfg.Service.Port)
	log.Infof("Pharmacist Service is starting on port: %d", cfg.Service.Port)

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

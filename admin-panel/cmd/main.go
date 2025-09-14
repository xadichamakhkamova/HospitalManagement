package main

import (
	app "admin-panel/internal/app"
	config "admin-panel/internal/config"
	"admin-panel/internal/repository"
	pq "admin-panel/internal/repository/postgres"
	service "admin-panel/internal/service"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"admin-panel/logger" 
)

func main() {
	// Logger init
	log := logger.NewLogger()

	// Config load
	cfg, err := config.Load("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	log.Info("Configuration loaded successfully")

	// Connect DB
	db, err := pq.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	log.Info("Connected to Postgres database")

	// Repository init
	queries := repository.NewAdminSqlc(db, log)	
	repo := repository.NewIAdminRepository(queries, log)
	log.Info("Repository initialized")

	// Service init
	srv := service.NewAdminService(repo)
	log.Info("Admin service initialized")

	// App init
	application := app.New(*srv, cfg.Service.Port)
	log.Infof("Admin Panel server is starting on port: %d", cfg.Service.Port)

	// Run server in goroutine
	go func() {
		application.MustRun()
	}()
	log.Info("Server is running...")

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop
	log.Info("Received signal: ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	application.Stop()
	log.Info("Server stopped gracefully")

	<-ctx.Done()
	log.Info("Shutdown complete")
}

package main

import (
	app "doctor-service/internal/app"
	config "doctor-service/internal/config"
	"doctor-service/internal/repository"
	pq "doctor-service/internal/repository/postgres"
	service "doctor-service/internal/service"

	adminservice "doctor-service/internal/clients/adminservice-client"
	patientservice "doctor-service/internal/clients/patientservice-client"

	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"doctor-service/logger" 
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
	queries := repository.NewDoctorSqlc(db, log)
	repo := repository.NewIDoctorRepository(queries, log)
	log.Info("Doctor repository initialized")

	// Connect to Admin Service
	conn1, err := adminservice.DialWithAdminService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Admin Service:", err)
	}
	log.Info("Connected to Admin Service")

	// Connect to Patient Service
	conn2, err := patientservice.DialWithPatientService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Patient Service:", err)
	}
	log.Info("Connected to Patient Service")

	// Service init
	srv := service.NewDoctorService(repo, conn1, conn2)
	log.Info("Doctor service initialized")

	// App init
	application := app.New(*srv, cfg.Service.Port)
	log.Infof("Doctor Service is starting on port: %d", cfg.Service.Port)
	
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

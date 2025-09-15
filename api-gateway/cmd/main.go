package main

import (
	_ "api-gateway/docs"
	config "api-gateway/internal/config"
	api "api-gateway/internal/https"

	adminService "api-gateway/internal/clients/adminclient-service"
	doctorService "api-gateway/internal/clients/doctorclient-service"
	nurseService "api-gateway/internal/clients/nurseclient-service"
	patientService "api-gateway/internal/clients/patientclient-service"
	pharmacistService "api-gateway/internal/clients/pharmacistclient-service"

	service "api-gateway/internal/service"
	"api-gateway/logger"

	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log := logger.NewLogger()

	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	log.Info("Configuration loaded successfully")

	conn1, err := adminService.DialWithAdminService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Admin Service:", err)
	}
	log.Info("Connected to Admin Service")

	conn2, err := doctorService.DialWithDoctorService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Doctor Service:", err)
	}
	log.Info("Connected to Doctor Service")

	conn3, err := nurseService.DialWithNurseService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Nurse Service:", err)
	}
	log.Info("Connected to Nurse Service")

	conn4, err := patientService.DialWithPatientService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Patient Service:", err)
	}
	log.Info("Connected to Patient Service")

	conn5, err := pharmacistService.DialWithPharmacistService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Pharmacist Service:", err)
	}
	log.Info("Connected to Pharmacist Service")

	clientService := service.NewServiceRepositoryClient(conn1, conn2, conn3, conn4, conn5)
	log.Info("Service clients initialized")

	srv := api.NewGin(clientService, cfg.ApiGateway.Port, log)
	addr := fmt.Sprintf(":%d", cfg.ApiGateway.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Starting API Gateway on: ", addr)
		if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			log.Fatal(err)
		}
	}()
	log.Info("Starting API Gateway on address:", addr)

	signalReceived := <-sigChan
	log.Info("Received signal:", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server shutdown error: ", err)
	}
	log.Info("Graceful shutdown complete.")
}

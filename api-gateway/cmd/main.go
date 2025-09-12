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

	// Konfiguratsiya faylini yuklash
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	log.Info("Configuration loaded successfully")

	// Admin Service bilan ulanish
	conn1, err := adminService.DialWithAdminService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Admin Service:", err)
	}
	log.Info("Connected to Admin Service")

	// Doctor Service bilan ulanish
	conn2, err := doctorService.DialWithDoctorService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Doctor Service:", err)
	}
	log.Info("Connected to Doctor Service")

	// Nurse Service bilan ulanish
	conn3, err := nurseService.DialWithNurseService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Nurse Service:", err)
	}
	log.Info("Connected to Nurse Service")

	// Patient Service bilan ulanish
	conn4, err := patientService.DialWithPatientService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Patient Service:", err)
	}
	log.Info("Connected to Patient Service")

	// Pharmacist Service bilan ulanish
	conn5, err := pharmacistService.DialWithPharmacistService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to Pharmacist Service:", err)
	}
	log.Info("Connected to Pharmacist Service")

	clientService := service.NewServiceRepositoryClient(conn1, conn2, conn3, conn4, conn5)
	log.Info("Service clients initialized")

	// API Gateway serverini ishga tushirish
	srv := api.NewGin(clientService, cfg.ApiGateway.Port)
	addr := fmt.Sprintf(":%d", cfg.ApiGateway.Port)

	// Signalni kutish uchun kanal yaratish (SIGINT yoki SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Starting API Gateway on: ", addr)
		if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			log.Fatal(err)
		}
	}()
	log.Info("Starting API Gateway on address:", addr)

	// Signalni qabul qilish
	signalReceived := <-sigChan
	log.Info("Received signal:", signalReceived)

	// Xizmatni to'xtatish uchun kontekst yaratish
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownRelease()

	// API Gatewayni to'xtatish
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server shutdown error: ", err)
	}
	log.Info("Graceful shutdown complete.")
}

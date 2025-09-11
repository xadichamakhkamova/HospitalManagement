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
	"fmt"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.Load("./config/config.yaml")
	if err != nil {
	}

	db, err := pq.ConnectDB(cfg)
	if err != nil {
	}

	queries := repository.NewDoctorSqlc(db)
	repo := repository.NewIDoctorRepository(queries)

	conn1, err := adminservice.DialWithAdminService(*cfg)
	if err != nil {
	}
	fmt.Print("Connected to Product Service")

	conn2, err := patientservice.DialWithPatientService(*cfg)
	if err != nil {
	}
	fmt.Print("Connected to Product Service")

	srv := service.NewDoctorService(repo, conn1, conn2)

	application := app.New(*srv, cfg.Service.Port)

	go func() {
		application.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop

	fmt.Println("Received signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	application.Stop()
	<-ctx.Done()
}

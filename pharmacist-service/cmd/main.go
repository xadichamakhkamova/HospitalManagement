package main

import (
	app "pharmacist-service/internal/app"
	config "pharmacist-service/internal/config"
	"pharmacist-service/internal/repository"
	pq "pharmacist-service/internal/repository/postgres"
	service "pharmacist-service/internal/service"
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

	queries := repository.NewPharmaSqlc(db)
	repo := repository.NewIPharmacistRepository(queries)

	srv := service.NewPharmaService(repo)

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

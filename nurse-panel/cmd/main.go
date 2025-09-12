package main

import (
	"context"
	"fmt"
	app "nurse-service/internal/app"
	config "nurse-service/internal/config"
	"nurse-service/internal/repository"
	pq "nurse-service/internal/repository/postgres"
	service "nurse-service/internal/service"

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

	queries := repository.NewNurseSqlc(db)
	repo := repository.NewINurseRepository(queries)

	srv := service.NewNurseService(repo)

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

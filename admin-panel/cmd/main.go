package main

import (
	app "admin-panel/internal/app"
	config "admin-panel/internal/config"
	"admin-panel/internal/repository"
	pq "admin-panel/internal/repository/postgres"
	service "admin-panel/internal/service"
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

	queries := repository.NewAdminSqlc(db)
	repo := repository.NewIAdminRepository(queries)

	srv := service.NewAdminService(repo)

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

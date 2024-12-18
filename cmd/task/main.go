package main

import (
	"context"
	"github.com/Ablyamitov/task/internal/config"
	"github.com/Ablyamitov/task/internal/handler"
	fiberserver "github.com/Ablyamitov/task/internal/server"
	"github.com/Ablyamitov/task/internal/storage/db/postgres"
	"github.com/Ablyamitov/task/internal/storage/repository"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := postgres.Connect(conf.DB.URL)
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}

	userRepository := repository.NewUserRepository(db)
	authHandler := handler.NewAuthHandler(userRepository, conf.App.Secret)
	adminHandler := handler.NewAdminHandler(userRepository)

	taskServer := fiberserver.NewServer(conf.Host, conf.Port, conf.App.Secret, authHandler, adminHandler)
	taskServer.Run()
	waitForShutdown(taskServer)
}

func waitForShutdown(taskServer fiberserver.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	taskServer.Stop(ctx)

}

package main

import (
	"context"
	"github.com/Ablyamitov/task/internal/config"
	fiberserver "github.com/Ablyamitov/task/internal/server"
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

	taskServer := fiberserver.NewServer(conf.Host, conf.Port)
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

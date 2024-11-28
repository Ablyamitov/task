package server

import (
	"context"
	"fmt"
	"github.com/Ablyamitov/task/internal/handler"
	"github.com/Ablyamitov/task/internal/middleware"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"time"
)

type Server interface {
	Run()
	Stop(ctx context.Context)
}

type TaskServer struct {
	app  *fiber.App
	host string
	port int
}

func NewServer(host string, port int, secret string, authHandler handler.AuthHandler, adminHandler handler.AdminHandler) Server {
	app := fiber.New()
	app.Use(fiberLogger.New())

	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	app.Get("/users", middleware.IsAdmin(secret), adminHandler.Users)

	server := &TaskServer{
		app:  app,
		host: host,
		port: port,
	}
	return server
}

func (server *TaskServer) Run() {
	go func() {
		addr := fmt.Sprintf("%s:%d", server.host, server.port)
		if err := server.app.Listen(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

}

func (server *TaskServer) Stop(ctx context.Context) {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := server.app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown server: %v", err)
	}
	log.Println("Server stopped gracefully")
}

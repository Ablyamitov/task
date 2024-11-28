package handler

import (
	"context"
	"fmt"
	"github.com/Ablyamitov/task/internal/storage/repository"
	"github.com/Ablyamitov/task/internal/web/mapper"
	"github.com/Ablyamitov/task/internal/web/response"
	"github.com/gofiber/fiber/v2"
	"log"
)

type AdminHandler interface {
	Users(c *fiber.Ctx) error
}

type adminHandler struct {
	UserRepository repository.UserRepository
}

func NewAdminHandler(userRepository repository.UserRepository) AdminHandler {
	return &adminHandler{UserRepository: userRepository}
}

func (adminHandler *adminHandler) Users(c *fiber.Ctx) error {
	method := "AdminHandler:Users"
	var Errors []string

	users, err := adminHandler.UserRepository.GetAll(context.Background())
	if err != nil {
		log.Printf(fmt.Sprintf("%s: Failed to fetch users: %v", method, err))
		Errors = append(Errors, "failed to fetch users")
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.UsersResponse{
				Status: false,
				Errors: &Errors,
			})
	}

	var usersDTO = make([]*response.UserDTO, len(users))
	for i := 0; i < len(users); i++ {
		usersDTO[i] = mapper.MapUserToUserDTO(&users[i])
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.UsersResponse{
			Data:   &response.Fetched{Users: usersDTO},
			Status: true,
		})
}

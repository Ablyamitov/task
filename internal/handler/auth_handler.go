package handler

import (
	"context"
	"github.com/Ablyamitov/task/internal/storage/repository"
	"github.com/Ablyamitov/task/internal/web/dto"
	"github.com/Ablyamitov/task/internal/web/mapper"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Register(c *fiber.Ctx) error
}

type authHandler struct {
	UserRepository repository.UserRepository
	Secret         string
}

func NewAuthHandler(userRepository repository.UserRepository) AuthHandler {
	return &authHandler{UserRepository: userRepository}
}

func (authHandler *authHandler) Register(c *fiber.Ctx) error {

	var userDTO dto.UserDTO

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user := mapper.MapUserDTOToUser(&userDTO)

	if err := authHandler.UserRepository.Create(context.Background(), user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create movie"})
	}

	response := mapper.MapUserToUserDTO(user)
	return c.Status(fiber.StatusCreated).JSON(response)

}

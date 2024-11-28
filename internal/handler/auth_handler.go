package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ablyamitov/task/internal/storage/repository"
	"github.com/Ablyamitov/task/internal/web/mapper"
	"github.com/Ablyamitov/task/internal/web/response"
	"github.com/gofiber/fiber/v2"
	"log"
	"regexp"
	"strings"
	"time"
)

type AuthHandler interface {
	Register(c *fiber.Ctx) error
}

type authHandler struct {
	UserRepository repository.UserRepository
}

func NewAuthHandler(userRepository repository.UserRepository) AuthHandler {
	return &authHandler{UserRepository: userRepository}
}

func (authHandler *authHandler) Register(c *fiber.Ctx) error {
	method := "AuthHandler:Register"
	var userDTO response.UserDTO
	var Errors []string

	if err := c.BodyParser(&userDTO); err != nil {

		log.Printf(fmt.Sprintf("%s: Invalid input: %v", method, err))
		Errors = append(Errors, "invalid input")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.SavedResponse{
				Status: false,
				Errors: &Errors,
			})
	}

	Errors = validateUserDTO(&userDTO)
	if len(Errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.SavedResponse{
				Status: false,
				Errors: &Errors,
			})
	}

	user := mapper.MapUserDTOToUser(&userDTO)

	if err := authHandler.UserRepository.Create(context.Background(), user); err != nil {
		log.Printf(fmt.Sprintf("%s: Failed to create user: %v", method, err))
		Errors = append(Errors, "failed to create user")
		if errors.Is(err, repository.ErrUserAlreadyExist) {
			Errors = append(Errors, repository.ErrUserAlreadyExist.Error())
			return c.Status(fiber.StatusConflict).JSON(
				response.SavedResponse{
					Status: false,
					Errors: &Errors,
				})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.SavedResponse{
				Status: false,
				Errors: &Errors,
			})

	}

	return c.Status(fiber.StatusCreated).JSON(
		response.SavedResponse{
			Data:   &response.Saved{Status: true},
			Status: true,
		})

}

func validateUserDTO(userDTO *response.UserDTO) []string {
	var Errors []string

	if len([]rune(userDTO.LastName)) < 2 {
		Errors = append(Errors, "LastName must be at least 2 char")
	}

	if len([]rune(userDTO.FirstName)) < 2 {
		Errors = append(Errors, "FirstName must be at least 2 char")
	}

	if strings.ToLower(userDTO.Gender) != "male" && strings.ToLower(userDTO.Gender) != "female" {
		Errors = append(Errors, "Gender must be male or female")
	}

	if _, err := time.Parse("02-01-2006", userDTO.BirthDate); err != nil {
		Errors = append(Errors, "BirthDate must be in the format 'DD-MM-YYYY'")
	}

	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	if !re.MatchString(userDTO.Phone) {
		Errors = append(Errors, "Phone must be a valid number")
	}

	return Errors
}

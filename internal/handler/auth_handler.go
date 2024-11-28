package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ablyamitov/task/internal/storage/model"
	"github.com/Ablyamitov/task/internal/storage/repository"
	"github.com/Ablyamitov/task/internal/web/mapper"
	"github.com/Ablyamitov/task/internal/web/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"regexp"
	"strings"
	"time"
)

type AuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	UserRepository repository.UserRepository
	Secret         string
}

func NewAuthHandler(userRepository repository.UserRepository, secret string) AuthHandler {
	return &authHandler{UserRepository: userRepository, Secret: secret}
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

	user.Role = "Role_User"
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

func (authHandler *authHandler) Login(c *fiber.Ctx) error {
	method := "AuthHandler:Login"
	type LoginRequest struct {
		Phone string `json:"phone"`
	}
	var Errors []string
	var loginRequest LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {

		log.Printf(fmt.Sprintf("%s: Invalid input: %v", method, err))
		Errors = append(Errors, "invalid input")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.SavedResponse{
				Status: false,
				Errors: &Errors,
			})
	}

	if !validatePhone(loginRequest.Phone) {
		log.Printf(fmt.Sprintf("%s: Phone must be a valid number", method))
		Errors = append(Errors, "Phone must be a valid number")
		return c.Status(fiber.StatusBadRequest).JSON(
			response.SavedResponse{
				Status: false,
				Errors: &Errors,
			})
	}

	var existingUser *model.User
	existingUser, err := authHandler.UserRepository.GetByPhone(context.Background(), loginRequest.Phone)
	if err != nil {
		log.Printf(fmt.Sprintf("%s: User with same phone does not exist: %v", method, err))
		Errors = append(Errors, "User with same phone does not exist")
		return c.Status(fiber.StatusNotFound).JSON(
			response.SavedResponse{
				Status: false,
				Errors: &Errors,
			})
	}

	claims := jwt.MapClaims{
		"id":   existingUser.ID,
		"role": existingUser.Role,
		"exp":  time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(authHandler.Secret))
	if err != nil {
		log.Printf("%s: Failed to generate token: %v", method, err)
		Errors = append(Errors, "Failed to generate token")
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.SavedResponse{
				Status: false,
				Errors: &Errors,
			})
	}

	c.Set("Authorization", "Bearer "+tokenString)

	return c.Status(fiber.StatusOK).JSON(
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

	if !validatePhone(userDTO.Phone) {
		Errors = append(Errors, "Phone must be a valid number")
	}

	return Errors
}

func validatePhone(phone string) bool {
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}

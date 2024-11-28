package middleware

import (
	"github.com/Ablyamitov/task/internal/web/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
)

func IsAdmin(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		method := "Middleware:IsAdmin"
		var Errors []string
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			log.Printf("%s: Authorization token is miss", method)
			Errors = append(Errors, "Authorization token is miss")
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.AccessResponse{
					Status: false,
					Errors: &Errors,
				})
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			log.Printf("%s: Invalid token: %v", method, err)

			Errors = append(Errors, "invalid token")
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.AccessResponse{
					Status: false,
					Errors: &Errors,
				})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Printf("%s: Invalid token claims", method)
			Errors = append(Errors, "Invalid token claims")
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.AccessResponse{
					Status: false,
					Errors: &Errors,
				})
		}

		role, ok := claims["role"]
		if !ok || role != "Role_Admin" {
			log.Printf("%s: Access denied", method)
			Errors = append(Errors, "Access denied")
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.AccessResponse{
					Status: false,
					Errors: &Errors,
				})
		}

		return c.Next()
	}
}

// Package middleware provides HTTP middleware for authentication and other purposes.
package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/utils"
	"github.com/spf13/viper"
)

const userIDKey = "user_id"

// NewAuth creates a new authentication middleware handler.
// It takes a Viper configuration object as a parameter and returns a Fiber handler.
func NewAuth(cfg *viper.Viper) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get token from header
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// check if token has "Bearer " prefix
		if len(token) < 7 || token[:7] != "Bearer " {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		// remove "Bearer " prefix
		token = token[7:]
		if token == "" {
			log.Println("Token validation failed")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// validate token
		claims, err := utils.ValidateToken(token, cfg)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		// if token is valid, set user to context
		c.Locals(userIDKey, claims.ID)
		return c.Next()
	}
}

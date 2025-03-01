// Package middleware provides HTTP middleware for authentication and other purposes.
package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const userIDKey = "user_id"

// NewAuth creates a new authentication middleware handler.
// It takes a Viper configuration object as a parameter and returns a Fiber handler.
func NewAuth(cfg *viper.Viper, log *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		// Log user authentication
		log.WithFields(logrus.Fields{
			"ip": ip,
		}).Info("Authenticating user")

		// Get token from header
		token := c.Get("Authorization")
		if token == "" {
			// log error
			log.WithFields(logrus.Fields{
				"ip":  ip,
				"err": "no token provided",
			}).Warn("Unauthorized access attempt")

			// return unauthorized
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Check if token has "Bearer " prefix
		if len(token) < 7 || token[:7] != "Bearer " {
			// log error
			log.WithFields(logrus.Fields{
				"ip":  ip,
				"err": "invalid token format",
			}).Warn("Unauthorized access attempt")

			// return unauthorized
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Remove "Bearer " prefix
		token = token[7:]
		if token == "" {
			// log error
			log.WithFields(logrus.Fields{
				"ip": ip,
			}).Warn("Unauthorized access attempt")

			// return unauthorized
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		// Validate token
		claims, err := utils.ValidateToken(token, cfg)
		if err != nil {
			// log error
			log.WithFields(logrus.Fields{
				"ip": ip,
			}).Errorf("Token validation failed")

			// return unauthorized
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Log user authentication
		log.WithFields(logrus.Fields{
			"ip":      ip,
			"user_id": claims.ID,
		}).Info("User authenticated")

		// If token is valid, set user to context
		c.Locals(userIDKey, claims.ID)
		return c.Next()
	}
}

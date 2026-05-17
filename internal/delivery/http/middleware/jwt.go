package middleware

import (
	"simpleorder/pkg/response"
	"simpleorder/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Protected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, fiber.StatusUnauthorized, "Missing Authorization header", nil)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid Authorization header format", nil)
		}

		token := parts[1]
		claims, err := utils.ValidateToken(token, secret)
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid or expired token", nil)
		}

		// Set user context
		c.Locals("user_id", claims.ID)
		c.Locals("user_role", claims.Role)
		c.Locals("username", claims.Username)

		return c.Next()
	}
}

func RoleAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("user_role")
		if role != "admin" {
			return response.Error(c, fiber.StatusForbidden, "Forbidden: Admin access required", nil)
		}
		return c.Next()
	}
}

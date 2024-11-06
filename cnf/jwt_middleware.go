package cnf

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "No token provided"})
	}

	// Format: "Bearer <token>"
	if len(token) < 7 || token[:7] != "Bearer " {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}
	token = token[7:] // Remove "Bearer "

	// Parse token
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized, invalid JWT"})
	}

	// Set user ID in context (optional)
	c.Locals("userID", claims.Subject) // Jika Anda menggunakan Subject untuk user ID

	return c.Next()
}

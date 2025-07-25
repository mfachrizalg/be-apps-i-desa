package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// JWTAuth verifies that the request contains a valid JWT token
func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get cookie from request
		cookie := c.Cookies("AppsIDesaCookie")
		// Check if the cookie is empty
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: No token provided",
			})
		}

		// Parse and validate the token
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			// Make sure the token method is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
				"error":   err.Error(),
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Token is not valid",
			})
		}

		return c.Next()
	}
}

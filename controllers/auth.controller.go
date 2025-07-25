package controllers

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *services.AuthService
	validate    *validator.Validate
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
		validate:    validator.New(),
	}
}

// Login handles user login requests
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var request dtos.LoginRequest

	// Parse request body
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate request
	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	// Process login
	response, err := c.authService.Login(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Login failed",
			"error":   err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "AppsIDesaCookie",
		Value:    response.Token,
		MaxAge:   3600,  // 1 hour
		Secure:   false, // Set to true if using HTTPS
		HTTPOnly: true,
		SameSite: "Lax",
	}

	ctx.Cookie(&cookie)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Logout handles user logout requests
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	response := c.authService.Logout()
	ctx.ClearCookie("AppsIDesaCookie")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

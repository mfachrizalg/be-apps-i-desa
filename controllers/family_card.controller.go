package controllers

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FamilyCardController struct {
	familyCardService *services.FamilyCardService
	validate          *validator.Validate
}

func NewFamilyCardController(familyCardService *services.FamilyCardService) *FamilyCardController {
	return &FamilyCardController{
		familyCardService: familyCardService,
		validate:          validator.New(),
	}
}

// AddFamilyCard handles the creation of a new family card
func (c *FamilyCardController) AddFamilyCard(ctx *fiber.Ctx) error {
	var request dtos.AddFamilyCardRequest

	// Parse the request body into the AddFamilyCardRequest struct
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate the request data
	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	// Call the service to create the family card
	response, err := c.familyCardService.CreateFamilyCard(&request, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create family card",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetFamilyCardByID retrieves a family card by its ID
func (c *FamilyCardController) GetFamilyCardByID(ctx *fiber.Ctx) error {
	familyCardID := ctx.Params("id")
	if familyCardID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Family card ID is required",
		})
	}

	response, err := c.familyCardService.GetFamilyCardByID(familyCardID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve family card",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

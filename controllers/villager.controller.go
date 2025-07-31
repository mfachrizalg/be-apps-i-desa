package controllers

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type VillagerController struct {
	villagerService *services.VillagerService
	validate        *validator.Validate
}

func NewVillagerController(villagerService *services.VillagerService) *VillagerController {
	return &VillagerController{
		villagerService: villagerService,
		validate:        validator.New(),
	}
}

func (c *VillagerController) CreateVillager(ctx *fiber.Ctx) error {
	var request dtos.AddVillagerRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Invalid request body",
			"Error":   err.Error(),
		})
	}

	if err := c.validate.Struct(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Validation failed",
			"Error":   err.Error(),
		})
	}

	response, err := c.villagerService.CreateVillager(&request, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Failed to create villager",
			"Error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

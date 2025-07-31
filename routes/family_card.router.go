package routes

import (
	"Apps-I_Desa_Backend/controllers"
	"Apps-I_Desa_Backend/middleware"
	"Apps-I_Desa_Backend/repositories"
	"Apps-I_Desa_Backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupFamilyCardRoutes(app *fiber.App) {
	familyCardRepo := repositories.NewFamilyCardRepository()
	familyCardService := services.NewFamilyCardService(familyCardRepo)
	familyCardController := controllers.NewFamilyCardController(familyCardService)

	api := app.Group("/api/family-cards")
	// Apply JWT middleware to all family card routes
	api.Use(middleware.JWTAuth())

	api.Post("/", familyCardController.AddFamilyCard)
	api.Get("/:id", familyCardController.GetFamilyCardByID)
}

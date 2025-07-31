package routes

import (
	"Apps-I_Desa_Backend/controllers"
	"Apps-I_Desa_Backend/middleware"
	"Apps-I_Desa_Backend/repositories"
	"Apps-I_Desa_Backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupVillageRoutes(app *fiber.App) {
	villageRepo := repositories.NewVillageRepository()
	villageService := services.NewVillageService(villageRepo)
	villageController := controllers.NewVillageController(villageService)

	api := app.Group("/api/villages")
	// Apply JWT middleware to all village routes
	api.Use(middleware.JWTAuth())

	api.Post("/", villageController.CreateVillage)
}

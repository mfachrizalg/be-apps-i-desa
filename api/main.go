package handler

import (
	"net/http"
	"sync"

	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/routes"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	fiberOnce    sync.Once
	fiberApp     *fiber.App
	httpAdaptor  http.Handler
	initAppError error
)

func setupRoutes(app *fiber.App) {
	routes.SetupUserRoutes(app)
	routes.SetupAuthRoutes(app)
	routes.SetupSubDimensionRoutes(app)
	routes.SetupVillagerRoutes(app)
	routes.SetupVillageRoutes(app)
	routes.SetupFamilyCardRoutes(app)
	routes.SetupDashboardRoutes(app)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Apps-I Desa API!")
	})
}

func initFiberApp() (*fiber.App, http.Handler, error) {
	fiberOnce.Do(func() {
		config.ConnectDB()

		fiberApp = fiber.New(fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			},
		})

		fiberApp.Use(logger.New())
		fiberApp.Use(recover.New())
		fiberApp.Use(cors.New())

		setupRoutes(fiberApp)
		httpAdaptor = adaptor.FiberApp(fiberApp)
	})

	return fiberApp, httpAdaptor, initAppError
}

func Handler(w http.ResponseWriter, r *http.Request) {
	_, handler, err := initFiberApp()
	if err != nil || handler == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	handler.ServeHTTP(w, r)
}

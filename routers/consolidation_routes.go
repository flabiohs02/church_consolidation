package routers

import (
	"church_consolidation/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupConsolidationRoutes(app *fiber.App, consolidationHandler *handler.ConsolidationHandler) {
	v1 := app.Group("/api/v1")
	{
		v1.Post("/consolidations", consolidationHandler.CreateConsolidation)
		v1.Get("/consolidations/:id", consolidationHandler.GetConsolidationByID)
		v1.Get("/consolidations", consolidationHandler.GetAllConsolidations)
		v1.Put("/consolidations/:id", consolidationHandler.UpdateConsolidation)
		v1.Delete("/consolidations/:id", consolidationHandler.DeleteConsolidation)
	}
}

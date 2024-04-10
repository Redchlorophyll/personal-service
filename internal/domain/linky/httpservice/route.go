package httpservice

import "github.com/gofiber/fiber/v2"

func (handler *LinkyHandler) SetRoute(app *fiber.App) {
	group := app.Group("/api/v1/linky")

	group.Get("/link", handler.GetLinky)
	group.Post("/link", handler.CreateLinky)
	// group.POST("/link/identifier")
	// group.Patch("/link/:linkId")
	// group.Delete("/link/:linkId")

	// group.Post("/ping")
}

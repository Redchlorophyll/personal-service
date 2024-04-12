package httpservice

import "github.com/gofiber/fiber/v2"

func (handler *LinkyHandler) SetRoute(app *fiber.App) {
	group := app.Group("/api/v1/linky")

	group.Get("/link", handler.GetLinky)
	group.Post("/link", handler.CreateLinky)
	group.Post("/link/identifier", handler.CreateIdentifier)
	// group.Patch("/link/:linkId", handler.DeleteLinky)
	group.Delete("/link/:linkId", handler.DeleteLinky)

	// group.Post("/ping")
}

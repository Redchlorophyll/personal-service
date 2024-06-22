package httpservice

import (
	utilsFunction "github.com/Redchlorophyll/personal-service/internal/utils/function"
	"github.com/gofiber/fiber/v2"
)

func (handler *LinkyHandler) SetRoute(app *fiber.App) {
	group := app.Group("/api/v1/linky")

	group.Get("/link", handler.GetLinky)
	group.Post("/link", utilsFunction.CheckUserAccess(), handler.CreateLinky)
	group.Post("/link/identifier", utilsFunction.CheckUserAccess(), handler.CreateIdentifier)
	group.Patch("/link/:linkId", utilsFunction.CheckUserAccess(), handler.UpdateLinky)
	group.Delete("/link/:linkId", utilsFunction.CheckUserAccess(), handler.DeleteLinky)
	group.Get("/profile/:fullNameSlug", handler.GetProfileLinky)

	// group.Post("/ping")
}

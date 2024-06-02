package httpservice

import "github.com/gofiber/fiber/v2"

func (handler *AccountHandler) SetRoute(app *fiber.App) {
	group := app.Group("/api/v1/account/")

	group.Get("/profile", handler.GetProfile)
	group.Post("/create", handler.CreateAccount)
	group.Post("/logout", handler.LogoutAccount)
	group.Post("/login", handler.Login)
}

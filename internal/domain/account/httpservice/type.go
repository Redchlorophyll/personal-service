package httpservice

import (
	"github.com/Redchlorophyll/personal-service/internal/domain/account/services"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountService services.AccountServiceProvider
}

type AccountHandlerConfig struct {
	AccountService services.AccountServiceProvider
}

type AccountHandlerProvider interface {
	SetRoute(app *fiber.App)
	GetProfile(fiberContext *fiber.Ctx) error
	CreateAccount(fiberContext *fiber.Ctx) error
	LogoutAccount(fiberContext *fiber.Ctx) error
	Login(fiberContext *fiber.Ctx) error
}

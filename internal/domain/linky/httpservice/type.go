package httpservice

import (
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/services"
	"github.com/gofiber/fiber/v2"
)

type LinkyHandler struct {
	LinkyService services.LinkyServiceProvider
}

type LinkyHandlerConfig struct {
	LinkyService services.LinkyServiceProvider
}

type LinkyHandlerProvider interface {
	SetRoute(app *fiber.App)
	GetLinky(fiberContext *fiber.Ctx) error
	DeleteLinky(fiberContext *fiber.Ctx) error
	CreateLinky(fiberContext *fiber.Ctx) error
	CreateIdentifier(fiberContext *fiber.Ctx) error
}

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
}

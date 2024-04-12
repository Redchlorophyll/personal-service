package httpservice

import (
	env "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/services"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
)

func (handler *LinkyHandler) CreateIdentifier(fiberContext *fiber.Ctx) error {
	token := fiberContext.Cookies("token")
	tokenFromEnv := env.GetEnvironmentVariables().SecretToken
	if token != tokenFromEnv {
		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		})
	}

	var request services.CreateIdentifierRequest

	err := fiberContext.BodyParser(&request)
	if err != nil {
		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: 400,
			Message:    "Bad Request!",
		})
	}

	result, err := handler.LinkyService.CreateIdentifier(fiberContext.Context(), request)
	if err != nil {
		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: 500,
			Message:    "Internal Server Error, Please try again later!",
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(result)
}

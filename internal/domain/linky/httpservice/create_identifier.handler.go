package httpservice

import (
	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *LinkyHandler) CreateIdentifier(fiberContext *fiber.Ctx) error {
	var request modelRequest.CreateIdentifierRequest

	err := fiberContext.BodyParser(&request)
	if err != nil {
		log.Error("[handler][CreateIdentifier] error when BodyParser(). ", err)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: 400,
			Message:    "Bad Request!",
		})
	}

	result, err := handler.LinkyService.CreateIdentifier(fiberContext.Context(), request)
	if err != nil {
		log.Error("[handler][CreateIdentifier] error when execute service in CreateIdentifier(). ", err, request)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: 500,
			Message:    "Internal Server Error, Please try again later!",
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(result)
}

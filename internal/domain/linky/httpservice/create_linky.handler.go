package httpservice

import (
	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *LinkyHandler) CreateLinky(fiberContext *fiber.Ctx) error {
	var request modelRequest.CreateLinkyRequest

	err := fiberContext.BodyParser(&request)
	if err != nil {
		log.Error("[handler][CreateLinky] error when BodyParser(). ", err)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		})
	}

	result, err := handler.LinkyService.CreateLinky(fiberContext.Context(), request)
	if err != nil {
		log.Error("[handler][CreateLinky] error when execute service CreateLinky(). ", err, request)

		return fiberContext.Status(fiber.StatusInternalServerError).JSON(utilsResponse.GeneralResponse{
			StatusCode: 500,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(result)
}

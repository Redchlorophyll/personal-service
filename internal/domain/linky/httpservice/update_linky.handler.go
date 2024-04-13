package httpservice

import (
	"strconv"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *LinkyHandler) UpdateLinky(fiberContext *fiber.Ctx) error {
	var request modelRequest.UpdateLinkyRequest

	err := fiberContext.BodyParser(&request)
	if err != nil {
		log.Error("[handler][UpdateLinky] error when BodyParser(). ", err)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		})
	}

	linkId, err := strconv.Atoi(fiberContext.Params("linkId"))
	if err != nil {
		log.Error("[handler][UpdateLinky] error when Params(). ", err)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		})
	}
	request.Id = linkId

	result, err := handler.LinkyService.UpdateLinky(fiberContext.Context(), request)
	if err != nil {
		log.Error("[handler][UpdateLinky] error when execute service UpdateLinky(). ", err, request)

		return fiberContext.Status(fiber.StatusInternalServerError).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(result)
}

package httpservice

import (
	"strconv"

	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *LinkyHandler) DeleteLinky(fiberContext *fiber.Ctx) error {
	request, err := strconv.Atoi(fiberContext.Params("linkId"))
	if err != nil {
		log.Error("[handler][DeleteLinky] error when Params(). ", err)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		})
	}

	result, err := handler.LinkyService.DeleteLinky(fiberContext.Context(), request)
	if err != nil {
		log.Error("[handler][DeleteLinky] error when execute service DeleteLinky(). ", err, request)

		return fiberContext.Status(fiber.StatusInternalServerError).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(result)
}

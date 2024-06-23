package httpservice

import (
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *AccountHandler) Login(fiberContext *fiber.Ctx) error {
	req := request.LoginRequest{}
	err := fiberContext.BodyParser(&req)
	context := fiberContext.Context()

	if err != nil {
		log.Error("[account][httpservice][Login] error when execute BodyParser(). ", context, err, req)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		})
	}

	err = handler.AccountService.Login(context, req)
	if err != nil {
		log.Error("[account][httpservice][Login] error when execute Login(). ", context, err, req)

		if err.Error() == "[ERROR]: not found" {
			return fiberContext.Status(fiber.StatusUnauthorized).JSON(utilsResponse.GeneralResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusUnauthorized].Error(),
			})
		}

		return fiberContext.Status(fiber.StatusInternalServerError).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(utilsResponse.GeneralResponse{
		StatusCode: fiber.StatusOK,
		Message:    "successfully login",
	})
}

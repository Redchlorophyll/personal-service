package httpservice

import (
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
)

func (handler *AccountHandler) Login(fiberContext *fiber.Ctx) error {
	req := request.LoginRequest{}
	err := fiberContext.BodyParser(&req)
	if err != nil {
		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		})
	}

	err = handler.AccountService.Login(fiberContext.Context(), req)
	if err != nil {
		return fiberContext.Status(fiber.StatusBadGateway).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadGateway,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadGateway].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(utilsResponse.GeneralResponse{
		StatusCode: fiber.StatusOK,
		Message:    "successfully create account",
	})
}

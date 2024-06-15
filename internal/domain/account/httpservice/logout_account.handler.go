package httpservice

import (
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *AccountHandler) LogoutAccount(fiberContext *fiber.Ctx) error {
	token := fiberContext.Cookies("token")
	context := fiberContext.Context()
	err := handler.AccountService.LogoutAccount(context, token)
	if err != nil {
		log.Error("[handler][CreateAccount] error when BodyParser(). ", context, err, token)

		if err.Error() == "[ERROR]: not found" {
			return fiberContext.Status(fiber.StatusUnauthorized).JSON(utilsResponse.GeneralResponse{
				StatusCode: fiber.StatusBadGateway,
				Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusUnauthorized].Error(),
			})
		}

		return fiberContext.Status(fiber.StatusInternalServerError).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(utilsResponse.GeneralResponse{
		StatusCode: fiber.StatusOK,
		Message:    "successfully logged out from your account",
	})
}

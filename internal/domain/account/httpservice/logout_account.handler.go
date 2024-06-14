package httpservice

import (
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *AccountHandler) LogoutAccount(fiberContext *fiber.Ctx) error {
	token := fiberContext.Cookies("token")

	err := handler.AccountService.LogoutAccount(fiberContext.Context(), token)
	if err != nil {
		log.Error("[handler][CreateAccount] error when BodyParser(). ", err)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(utilsResponse.GeneralResponse{
		StatusCode: fiber.StatusOK,
		Message:    "successfully logged out from your account",
	})
}

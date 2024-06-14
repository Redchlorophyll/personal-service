package httpservice

import (
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *AccountHandler) GetProfile(fiberContext *fiber.Ctx) error {
	token := fiberContext.Cookies("token")

	user, err := handler.AccountService.GetProfile(fiberContext.Context(), token)
	if err != nil {
		log.Error("[handler][CreateAccount] error when BodyParser(). ", err)

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(response.GetProfileResponse{
		GeneralResponse: utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		},
		Data: &user,
	})
}

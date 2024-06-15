package httpservice

import (
	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/log"
)

func (handler *AccountHandler) CreateAccount(fiberContext *fiber.Ctx) error {
	var request modelRequest.CreateAccountRequest

	file, err := fiberContext.FormFile("profile_img")
	request.ProfileImg = file
	if err != nil {
		request.ProfileImg = nil
	}

	request.UserName = fiberContext.FormValue("user_name")
	request.FullName = fiberContext.FormValue("full_name")
	request.Description = fiberContext.FormValue("description")
	request.Email = fiberContext.FormValue("email")
	request.Password = fiberContext.FormValue("password")

	context := fiberContext.Context()
	err = handler.AccountService.CreateAccount(context, request)

	if err == nil {
		// success
		return fiberContext.Status(fiber.StatusOK).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusOK,
			Message:    "successfully create account",
		})
	}

	if err.Error() == "[ERROR]: duplicate key" {
		log.Error("[account][httpservice][CreateAccount] error duplicate key. ", context, err, request)
		ErrorCode := 1000

		return fiberContext.Status(fiber.StatusBadRequest).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			ErrorCode:  &ErrorCode,
			Message:    "bad request, email or username already used",
		})
	}

	log.Error("[account][httpservice][CreateAccount] error when execute CreateAccount(). ", context, err, request)

	return fiberContext.Status(fiber.StatusBadGateway).JSON(utilsResponse.GeneralResponse{
		StatusCode: fiber.StatusBadGateway,
		Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadGateway].Error(),
	})
}

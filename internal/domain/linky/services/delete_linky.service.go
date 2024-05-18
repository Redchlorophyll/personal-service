package services

import (
	"context"

	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (service *LinkyService) DeleteLinky(context context.Context, request int) (utilsResponse.GeneralResponse, error) {
	err := service.LinkyRepository.SoftDeleteLinky(context, request)
	if err != nil {
		log.Error("[service][DeleteLinky] error when querying data in GetLinkyIdentifierById(). ", err, context, request)

		return utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		}, nil
	}

	return utilsResponse.GeneralResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success delete linky data!",
	}, nil
}

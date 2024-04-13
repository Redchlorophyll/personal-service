package services

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (service *LinkyService) CreateIdentifier(context context.Context, request request.CreateIdentifierRequest) (utilsResponse.GeneralResponse, error) {
	err := service.ContentIdentifierRepository.CreateIdentifier(
		context,
		request,
	)
	if err != nil {
		log.Error("[service][CreateLinky] error when execute sql query in CreateLinky(). ", err, context, request)

		return utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		}, nil
	}

	return utilsResponse.GeneralResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success create content identifier data",
	}, nil
}

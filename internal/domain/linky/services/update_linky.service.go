package services

import (
	"context"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (service *LinkyService) UpdateLinky(context context.Context, request modelRequest.UpdateLinkyRequest) (utilsResponse.GeneralResponse, error) {
	err := request.Validator()
	if err != nil {
		log.Error("[service][UpdateLinky] error when execute request.Validator(). ", err, context, request)

		return utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusBadRequest].Error(),
		}, nil
	}

	err = service.LinkyRepository.UpdateLinky(context, request)
	if err != nil {
		log.Error("[service][UpdateLinky] error when execute sql query in UpdateLinky(). ", err, context, request)

		return utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		}, nil
	}

	return utilsResponse.GeneralResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success update linky data",
	}, nil
}

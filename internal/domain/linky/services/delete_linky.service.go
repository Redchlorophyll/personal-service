package services

import (
	"context"

	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (service *LinkyService) DeleteLinky(context context.Context, request int) (utilsResponse.GeneralResponse, error) {
	err := service.LinkyRepository.SoftDeleteLinky(context, request)
	if err != nil {
		log.Error("[service][DeleteLinky] error when querying data in GetLinkyIdentifierById function. ", err, context, request)
		errorCode := 1
		return utilsResponse.GeneralResponse{
			ErrorCode:  &errorCode,
			StatusCode: 500,
			Message:    "Internal Server Error, Please try again later!",
		}, nil
	}

	return utilsResponse.GeneralResponse{
		StatusCode: 200,
		Message:    "Success delete linky data!",
	}, nil
}

package services

import (
	"context"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (service *LinkyService) UpdateLinky(context context.Context, request modelRequest.UpdateLinkyRequest) (utilsResponse.GeneralResponse, error) {
	err := service.LinkyRepository.UpdateLinky(context, request)
	if err != nil {
		log.Error("[service][UpdateLinky] error when execute sql query in UpdateLinky(). ", err, context, request)
		errorCode := 1
		return utilsResponse.GeneralResponse{
			ErrorCode:  &errorCode,
			StatusCode: 500,
			Message:    "Internal Server Error, Please try again later!",
		}, nil
	}

	return utilsResponse.GeneralResponse{
		StatusCode: 200,
		Message:    "Success update linky data!",
	}, nil
}

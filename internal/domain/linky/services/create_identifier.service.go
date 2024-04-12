package services

import (
	"context"

	contentIdentifierTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/content_identifier_table"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2/log"
)

type CreateIdentifierRequest struct {
	Identifier string `json:"identifier"`
	Title      string `json:"title"`
}

func (service *LinkyService) CreateIdentifier(context context.Context, request CreateIdentifierRequest) (utilsResponse.GeneralResponse, error) {
	err := service.ContentIdentifierRepository.CreateIdentifier(
		context,
		contentIdentifierTable.CreateIdentifierRequest(request),
	)
	if err != nil {
		log.Error("[service][CreateLinky] error when execute sql query in CreateLinky function. ", err, context, request)
		errorCode := 1
		return utilsResponse.GeneralResponse{
			ErrorCode:  &errorCode,
			StatusCode: 500,
			Message:    "Internal Server Error, Please try again later!",
		}, nil
	}

	return utilsResponse.GeneralResponse{
		StatusCode: 200,
		Message:    "Success create content identifier data!",
	}, nil
}

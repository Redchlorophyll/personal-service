package services

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/linky_table"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2/log"
)

type CreateLinkyRequest struct {
	IdentiferId int    `json:"identifier_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UrlAnchor   string `json:"url_anchhor"`
	ImageUrl    string `json:"img_url"`
}

func (service *LinkyService) CreateLinky(context context.Context, request CreateLinkyRequest) (utilsResponse.GeneralResponse, error) {
	_, err := service.LinkyRepository.GetLinkyIdentifierById(
		context,
		&request.IdentiferId,
	)
	if err != nil {
		log.Error("[service][CreateLinky] error when querying data in GetLinkyIdentifierById function. ", err, context, request)
		errorCode := 1
		return utilsResponse.GeneralResponse{
			ErrorCode:  &errorCode,
			StatusCode: 500,
			Message:    "Internal Server Error, Please try again later!",
		}, nil
	}

	// create linky data
	id, err := service.LinkyRepository.CreateLinky(
		context,
		linky_table.CreateLinkyRequest(request),
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

	// add join row for link & link identifier
	err = service.LinkyRepository.CreateLinkyIdentifier(context, linky_table.CreateLinkyIdentifierRequest{
		ContentIdentiferId: &request.IdentiferId,
		LinkId:             id,
	})
	if err != nil {
		log.Error("[service][CreateLinky] error when execute sql query in CreateLinkyIdentifier function. ", err, context, request)
		errorCode := 1
		return utilsResponse.GeneralResponse{
			ErrorCode:  &errorCode,
			StatusCode: 500,
			Message:    "Internal Server Error, Please try again later!",
		}, nil
	}

	return utilsResponse.GeneralResponse{
		StatusCode: 200,
		Message:    "Success create linky data!",
	}, nil
}
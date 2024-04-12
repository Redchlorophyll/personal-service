package services

import (
	"context"

	linkyTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/linky_table"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2/log"
)

type UpdateLinkyRequest struct {
	IdentiferId int    `json:"identifier_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UrlAnchor   string `json:"url_anchhor" db:"url_anchor"`
	ImageUrl    string `json:"img_url" db:"img_url"`
	Id          int
}

func (service *LinkyService) UpdateLinky(context context.Context, request UpdateLinkyRequest) (utilsResponse.GeneralResponse, error) {
	err := service.LinkyRepository.UpdateLinky(context, linkyTable.UpdateLinkyRequest(request))
	if err != nil {
		log.Error("[service][UpdateLinky] error when execute sql query in UpdateLinky function. ", err, context, request)
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

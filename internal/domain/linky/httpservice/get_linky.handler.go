package httpservice

import (
	"strconv"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsRequest "github.com/Redchlorophyll/personal-service/internal/utils/model/request"
	"github.com/Redchlorophyll/personal-service/internal/utils/model/response"

	"github.com/gofiber/fiber/v2"
)

func (handler *LinkyHandler) GetLinky(fiberContext *fiber.Ctx) error {
	linkType := fiberContext.Params("type")
	page, errParse := strconv.ParseInt(fiberContext.Params("page"), 10, 64)
	if errParse != nil {
		page = 1
	}
	perPage, errParse := strconv.ParseInt(fiberContext.Params("per_page"), 10, 64)
	if errParse != nil {
		perPage = 10
	}

	request := request.GetLinkyRequestQuery{
		Type: linkType,
		PaginationRequestQuery: utilsRequest.PaginationRequestQuery{
			Page:    int(page),
			PerPage: int(perPage),
		},
	}

	result, err := handler.LinkyService.GetLinky(fiberContext.Context(), request)

	if err != nil {
		return fiberContext.Status(fiber.StatusInternalServerError).JSON(response.GeneralResponse{
			StatusCode: 500,
			Message:    "there is something wrong in the system, please contact developer",
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(result)
}

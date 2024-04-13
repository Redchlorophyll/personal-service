package httpservice

import (
	"strconv"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsRequest "github.com/Redchlorophyll/personal-service/internal/utils/model/request"
	modelResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *LinkyHandler) GetLinky(fiberContext *fiber.Ctx) error {
	identifer := fiberContext.Query("identifier", "")
	page, errParse := strconv.ParseInt(fiberContext.Query("page"), 10, 64)
	if errParse != nil {
		page = 1
	}
	perPage, errParse := strconv.ParseInt(fiberContext.Query("per_page"), 10, 64)
	if errParse != nil {
		perPage = 10
	}

	request := modelRequest.GetLinkyRequestQuery{
		Identifier: identifer,
		PaginationRequestQuery: utilsRequest.PaginationRequestQuery{
			Page:    int(page),
			PerPage: int(perPage),
		},
	}

	result, err := handler.LinkyService.GetLinky(fiberContext.Context(), request)

	if err != nil {
		log.Error("[handler][GetLinky] error when execute GetLinky(). ", err, request)

		return fiberContext.Status(fiber.StatusInternalServerError).JSON(modelResponse.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(result)
}

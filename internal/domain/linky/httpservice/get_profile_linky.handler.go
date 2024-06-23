package httpservice

import (
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (handler *LinkyHandler) GetProfileLinky(fiberContext *fiber.Ctx) error {
	req := fiberContext.Params("fullNameSlug")
	result, err := handler.LinkyService.GetProfileLinky(fiberContext.Context(), request.GetProfileLinkyServiceRequest{
		UrlSlug: req,
	})
	if err != nil {
		log.Error("[linky][handler][GetProfileLinky] error when execute service GetProfileLinky(). ", err, req)

		return fiberContext.Status(fiber.StatusInternalServerError).JSON(utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
		})
	}

	return fiberContext.Status(fiber.StatusOK).JSON(result)
}

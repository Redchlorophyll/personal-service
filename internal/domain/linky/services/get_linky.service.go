package services

import (
	"context"

	requestModel "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	"github.com/Redchlorophyll/personal-service/internal/utils/function"
	utilsRequest "github.com/Redchlorophyll/personal-service/internal/utils/model/request"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (service *LinkyService) GetLinky(context context.Context, request requestModel.GetLinkyRequestQuery) (response.GetLinkyResponse, error) {
	linkData, err := service.LinkyRepository.GetLinky(context, requestModel.GetLinkyRequestQuery{
		PaginationRequestQuery: utilsRequest.PaginationRequestQuery{
			Page:    request.Page,
			PerPage: request.PerPage,
		},
		Identifier: request.Identifier,
	})
	if err != nil {
		log.Error("[service][GetLinky] error when querying data in getLinky(). ", err, context, request)

		return response.GetLinkyResponse{
			GeneralResponse: utilsResponse.GeneralResponse{
				StatusCode: fiber.StatusInternalServerError,
				Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
			},
			Pagination: nil,
			Metadata:   nil,
		}, nil
	}

	totalItem, err := service.LinkyRepository.GetTotalLinkyItem(context, request.Identifier)
	if err != nil {
		log.Error("[service][GetLinky] error when querying data in GetTotalLinkyItem(). ", err, context, request)

		return response.GetLinkyResponse{
			GeneralResponse: utilsResponse.GeneralResponse{
				StatusCode: fiber.StatusInternalServerError,
				Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
			},
			Pagination: nil,
			Metadata:   nil,
		}, nil
	}

	metadata, err := service.LinkyRepository.GetLinkyIdentifier(context, request.Identifier)
	if err != nil {
		log.Error("[service][GetLinky] error when querying data in GetLinkyIdentifier function. ", err, context, request)

		return response.GetLinkyResponse{
			GeneralResponse: utilsResponse.GeneralResponse{
				StatusCode: fiber.StatusInternalServerError,
				Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusInternalServerError].Error(),
			},
			Pagination: nil,
			Metadata:   nil,
		}, nil
	}

	pagination := function.GetPagination(utilsRequest.GetpaginationRequest{
		Page:      request.Page,
		PerPage:   request.PerPage,
		TotalItem: totalItem,
	})

	return response.GetLinkyResponse{
		GeneralResponse: utilsResponse.GeneralResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Success retrieve linky data!",
		},
		Data:       linkData,
		Pagination: &pagination,
		Metadata:   &metadata,
	}, nil
}

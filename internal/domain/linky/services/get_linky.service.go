package services

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	"github.com/Redchlorophyll/personal-service/internal/utils/function"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"

	utilsRequest "github.com/Redchlorophyll/personal-service/internal/utils/model/request"
)

func (service *LinkyService) GetLinky(context context.Context, request request.GetLinkyRequestQuery) (response.GetLinkyResponse, error) {
	linkData, err := service.LinkyRepository.GetLinky(context, utilsRequest.PaginationRequestQuery{
		Page:    request.Page,
		PerPage: request.PerPage,
	})

	if err != nil {
		return response.GetLinkyResponse{}, nil
	}

	totalItem, err := service.LinkyRepository.GetTotalLinkyItem(context)
	if err != nil {
		return response.GetLinkyResponse{
			GeneralResponse: utilsResponse.GeneralResponse{
				StatusCode: 500,
				Message:    "Internal Server Error!",
			},
		}, nil
	}

	offset := function.GetPaginationOffset(request.Page)

	return response.GetLinkyResponse{
		GeneralResponse: utilsResponse.GeneralResponse{
			StatusCode: 200,
			Message:    "Success!",
		},
		Data: linkData,
		Pagination: utilsResponse.PaginationResponse{
			Page:      request.Page,
			PerPage:   request.PerPage,
			FromItem:  offset,
			ToItem:    offset + 1 + request.PerPage,
			TotalItem: totalItem,
		},
	}, nil
}

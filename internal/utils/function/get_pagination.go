package function

import (
	"github.com/Redchlorophyll/personal-service/internal/utils/model/request"
	"github.com/Redchlorophyll/personal-service/internal/utils/model/response"
)

func GetPagination(pagination request.GetpaginationRequest) response.PaginationResponse {
	offset := GetPaginationOffset(pagination.Page)

	return response.PaginationResponse{
		Page:      pagination.Page,
		PerPage:   pagination.PerPage,
		FromItem:  offset + 1,
		ToItem:    offset + pagination.PerPage,
		TotalItem: pagination.TotalItem,
	}
}

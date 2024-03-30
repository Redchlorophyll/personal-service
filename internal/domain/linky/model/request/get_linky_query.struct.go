package request

import (
	utilsRequest "github.com/Redchlorophyll/personal-service/internal/utils/model/request"
)

type GetLinkyRequestQuery struct {
	utilsRequest.PaginationRequestQuery

	Identifier string `query:"identifier"`
}

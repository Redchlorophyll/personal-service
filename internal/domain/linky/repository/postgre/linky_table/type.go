package linky_table

import (
	"context"
	"database/sql"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	utilsRequest "github.com/Redchlorophyll/personal-service/internal/utils/model/request"
)

type LinkyTableRepository struct {
	Db *sql.DB
}

type LinkyTableRepositoryConfig struct {
	Db *sql.DB
}

type LinkyTableRepositoryProvider interface {
	GetLinky(context context.Context, pagination utilsRequest.PaginationRequestQuery) ([]response.LinkyItem, error)

	GetTotalLinkyItem(context context.Context) (int, error)
}

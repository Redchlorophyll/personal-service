package linky_table

import (
	"context"
	"database/sql"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
)

type LinkyTableRepository struct {
	Db *sql.DB
}

type LinkyTableRepositoryConfig struct {
	Db *sql.DB
}

type LinkyTableRepositoryProvider interface {
	GetLinky(context context.Context, request request.GetLinkyRequestQuery) ([]response.LinkyItem, error)

	GetTotalLinkyItem(context context.Context, request string) (int, error)

	GetLinkyIdentifier(context context.Context, request string) (response.GetLinkyIdentifierResponse, error)

	GetLinkyIdentifierById(context context.Context, request *int) (response.GetLinkyIdentifierResponse, error)

	CreateLinky(context context.Context, request CreateLinkyRequest) (*int, error)

	CreateLinkyIdentifier(context context.Context, request CreateLinkyIdentifierRequest) error

	SoftDeleteLinky(context context.Context, request int) error

	UpdateLinky(context context.Context, request UpdateLinkyRequest) error
}

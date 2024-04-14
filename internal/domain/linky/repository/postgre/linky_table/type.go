package linky_table

import (
	"context"
	"database/sql"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	modelResponse "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
)

type LinkyTableRepository struct {
	Db *sql.DB
}

type LinkyTableRepositoryConfig struct {
	Db *sql.DB
}

type LinkyTableRepositoryProvider interface {
	GetLinky(context context.Context, request modelRequest.GetLinkyRequestQuery) ([]modelResponse.LinkyItem, error)

	GetTotalLinkyItem(context context.Context, request string) (int, error)

	GetLinkyIdentifier(context context.Context, request string) (modelResponse.GetLinkyIdentifierResponse, error)

	GetLinkyIdentifierById(context context.Context, request *int) (modelResponse.GetLinkyIdentifierResponse, error)

	CreateLinky(context context.Context, request modelRequest.CreateLinkyRequest) (*int, error)

	CreateLinkyIdentifier(context context.Context, request modelRequest.CreateLinkyIdentifierRequest) error

	SoftDeleteLinky(context context.Context, request int) error

	UpdateLinky(context context.Context, request modelRequest.UpdateLinkyRequest) error
}

package account_url_slug_table

import (
	"context"
	"database/sql"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
)

type AccountUrlSlugTableRepositoryConfig struct {
	Db *sql.DB
}

type AccountUrlSlugTableRepository struct {
	Db *sql.DB
}

type AccountUrlSlugTableRepositoryProvider interface {
	GetAccountUrlSlug(context context.Context, request request.GetAccountUrlSlugRepositoryRequest) (response.GetAccountUrlSlugRepositoryResponse, error)
	CreateAccountUrlSlug(context context.Context, request request.CreateAccountUrlSlugRepositoryRequest) error
}

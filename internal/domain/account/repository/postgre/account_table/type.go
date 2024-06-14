package account_table

import (
	"context"
	"database/sql"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
)

type AccountTableRepository struct {
	Db *sql.DB
}

type AccountTableRepositoryConfig struct {
	Db *sql.DB
}

type AccountTableRepositoryProvider interface {
	CreateAccount(context context.Context, request request.CreateAccountRepositoryRequest) error
	GetAccountBySessionToken(context context.Context, request string) (response.GetAccountRespositoryResponse, error)
	RevokeSessionToken(context context.Context, request string) error
	SetSessionToken(context context.Context, request request.SetSessionTokenRequest) error
	GetAccountByEmail(context context.Context, request string) (response.GetAccountRespositoryResponse, error)
	GetAccountsByEmailOrUsername(context context.Context, request request.GetAccountByEmailOrUsernameRequest) ([]response.GetAccountRespositoryResponse, error)
}

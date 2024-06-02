package services

import (
	"context"

	AccountTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/account_table"
)

type AccountService struct {
	AccountRepository AccountTable.AccountTableRepositoryProvider
}

type AccountServiceConfig struct {
	AccountRepository AccountTable.AccountTableRepositoryProvider
}

type AccountServiceProvider interface {
	CreateAccount(context context.Context) error
	GetProfile(context context.Context) error
	Login(context context.Context) error
	LogoutAccount(context context.Context) error
}

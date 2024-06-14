package services

import (
	"context"
	"time"

	"github.com/Redchlorophyll/personal-service/internal/config/firebase"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	AccountTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/account_table"
)

type AccountService struct {
	AccountRepository AccountTable.AccountTableRepositoryProvider
	FirebaseService   firebase.FirebaseServiceProvider
}

type AccountServiceConfig struct {
	AccountRepository AccountTable.AccountTableRepositoryProvider
	FirebaseService   firebase.FirebaseServiceProvider
}

type AccountServiceProvider interface {
	CreateAccount(context context.Context, request request.CreateAccountRequest) error
	GetProfile(context context.Context, req string) (request.ProfileData, error)
	Login(context context.Context, request request.LoginRequest) error
	LogoutAccount(context context.Context, request string) error
	VerifyTokenExpiration(context context.Context, sessionTokenExpiredAt *time.Time) error
}

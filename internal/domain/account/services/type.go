package services

import (
	"context"
	"time"

	envVariable "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	"github.com/Redchlorophyll/personal-service/internal/config/firebase"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	AccountTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/account_table"
	AccountUrlSlugTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/account_url_slug_table"
	SocialMediaTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/social_media_table"
)

type AccountService struct {
	AccountRepository        AccountTable.AccountTableRepositoryProvider
	FirebaseService          firebase.FirebaseServiceProvider
	SocialMediaRepository    SocialMediaTable.SocialMediaTableRepositoryProvider
	AccountUrlSlugRepository AccountUrlSlugTable.AccountUrlSlugTableRepositoryProvider
	Env                      envVariable.Config
}

type AccountServiceConfig struct {
	AccountRepository        AccountTable.AccountTableRepositoryProvider
	FirebaseService          firebase.FirebaseServiceProvider
	SocialMediaRepository    SocialMediaTable.SocialMediaTableRepositoryProvider
	AccountUrlSlugRepository AccountUrlSlugTable.AccountUrlSlugTableRepositoryProvider
	Env                      envVariable.Config
}

type AccountServiceProvider interface {
	CreateAccount(context context.Context, request request.CreateAccountRequest) error
	GetProfile(context context.Context, req string) (request.ProfileData, error)
	Login(context context.Context, request request.LoginRequest) error
	LogoutAccount(context context.Context, request string) error
	VerifyTokenExpiration(context context.Context, sessionTokenExpiredAt *time.Time) error
}

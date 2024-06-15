package services

import (
	"context"

	envVariable "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	AccountService "github.com/Redchlorophyll/personal-service/internal/domain/account/services"
	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	modelResponse "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	contentIdentifierTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/content_identifier_table"
	linkyTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/linky_table"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
)

type LinkyService struct {
	LinkyRepository             linkyTable.LinkyTableRepositoryProvider
	ContentIdentifierRepository contentIdentifierTable.ContentIdentifierTableRepositoryProvider
	Env                         envVariable.Config
	AccountService              AccountService.AccountServiceProvider
}

type LinkyServiceConfig struct {
	LinkyRepository             linkyTable.LinkyTableRepositoryProvider
	ContentIdentifierRepository contentIdentifierTable.ContentIdentifierTableRepositoryProvider
	Env                         envVariable.Config
	AccountService              AccountService.AccountServiceProvider
}

type LinkyServiceProvider interface {
	GetLinky(context context.Context, request modelRequest.GetLinkyRequestQuery) (modelResponse.GetLinkyResponse, error)

	CreateLinky(context context.Context, request modelRequest.CreateLinkyRequest) (utilsResponse.GeneralResponse, error)

	CreateIdentifier(context context.Context, request modelRequest.CreateIdentifierRequest) (utilsResponse.GeneralResponse, error)

	DeleteLinky(context context.Context, request int) (utilsResponse.GeneralResponse, error)

	UpdateLinky(context context.Context, request modelRequest.UpdateLinkyRequest) (utilsResponse.GeneralResponse, error)
}

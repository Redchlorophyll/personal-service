package services

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	contentIdentifierTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/content_identifier_table"
	linkyTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/linky_table"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
)

type LinkyService struct {
	LinkyRepository             linkyTable.LinkyTableRepositoryProvider
	ContentIdentifierRepository contentIdentifierTable.ContentIdentifierTableRepositoryProvider
}

type LinkyServiceConfig struct {
	LinkyRepository             linkyTable.LinkyTableRepositoryProvider
	ContentIdentifierRepository contentIdentifierTable.ContentIdentifierTableRepositoryProvider
}

type LinkyServiceProvider interface {
	GetLinky(context context.Context, request request.GetLinkyRequestQuery) (response.GetLinkyResponse, error)

	CreateLinky(context context.Context, request CreateLinkyRequest) (utilsResponse.GeneralResponse, error)

	CreateIdentifier(context context.Context, request CreateIdentifierRequest) (utilsResponse.GeneralResponse, error)
}

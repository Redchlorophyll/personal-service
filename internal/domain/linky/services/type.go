package services

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	linkyTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/linky_table"
)

type LinkyService struct {
	LinkyRepository linkyTable.LinkyTableRepositoryProvider
}

type LinkyServiceConfig struct {
	LinkyRepository linkyTable.LinkyTableRepositoryProvider
}

type LinkyServiceProvider interface {
	GetLinky(context context.Context, request request.GetLinkyRequestQuery) (response.GetLinkyResponse, error)
}

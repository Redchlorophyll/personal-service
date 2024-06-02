package personal_service

import (
	AccountHandler "github.com/Redchlorophyll/personal-service/internal/domain/account/httpservice"
	AccountService "github.com/Redchlorophyll/personal-service/internal/domain/account/services"
	linkyHandler "github.com/Redchlorophyll/personal-service/internal/domain/linky/httpservice"
	LinkyService "github.com/Redchlorophyll/personal-service/internal/domain/linky/services"
)

type HTTPService struct {
	LinkyHandler   linkyHandler.LinkyHandlerProvider
	AccountHandler AccountHandler.AccountHandlerProvider
}

type Service struct {
	LinkyService   LinkyService.LinkyServiceProvider
	AccountService AccountService.AccountServiceProvider
}

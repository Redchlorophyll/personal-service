package personal_service

import (
	linkyHandler "github.com/Redchlorophyll/personal-service/internal/domain/linky/httpservice"
	LinkyService "github.com/Redchlorophyll/personal-service/internal/domain/linky/services"
)

type HTTPService struct {
	LinkyHandler linkyHandler.LinkyHandlerProvider
}

type Service struct {
	LinkyService LinkyService.LinkyServiceProvider
}

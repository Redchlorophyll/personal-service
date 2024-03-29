package personal_service

import (
	linkyHttpService "github.com/Redchlorophyll/personal-service/internal/domain/linky/httpservice"
)

func InitializeService(serv *Service) HTTPService {
	return HTTPService{
		LinkyHandler: linkyHttpService.NewHandler(linkyHttpService.LinkyHandlerConfig{
			LinkyService: serv.LinkyService,
		}),
	}
}

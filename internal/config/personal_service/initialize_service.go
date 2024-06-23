package personal_service

import (
	AccountHttpService "github.com/Redchlorophyll/personal-service/internal/domain/account/httpservice"
	linkyHttpService "github.com/Redchlorophyll/personal-service/internal/domain/linky/httpservice"
)

func InitializeService(serv *Service) HTTPService {
	return HTTPService{
		LinkyHandler: linkyHttpService.NewHandler(linkyHttpService.LinkyHandlerConfig{
			LinkyService: serv.LinkyService,
		}),
		AccountHandler: AccountHttpService.NewAccountHandler(AccountHttpService.AccountHandlerConfig{
			AccountService: serv.AccountService,
		}),
	}
}

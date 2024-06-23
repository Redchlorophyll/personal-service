package httpservice

func NewAccountHandler(config AccountHandlerConfig) AccountHandlerProvider {
	return &AccountHandler{
		AccountService: config.AccountService,
	}
}

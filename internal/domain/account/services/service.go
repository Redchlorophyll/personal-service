package services

func NewAccountService(config AccountServiceConfig) AccountServiceProvider {
	return &AccountService{}
}

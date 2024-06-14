package services

func NewAccountService(config AccountServiceConfig) AccountServiceProvider {
	return &AccountService{
		AccountRepository: config.AccountRepository,
		FirebaseService:   config.FirebaseService,
	}
}

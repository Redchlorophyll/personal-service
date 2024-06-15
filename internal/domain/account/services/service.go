package services

func NewAccountService(config AccountServiceConfig) AccountServiceProvider {
	return &AccountService{
		Env:                   config.Env,
		AccountRepository:     config.AccountRepository,
		SocialMediaRepository: config.SocialMediaRepository,
		FirebaseService:       config.FirebaseService,
	}
}

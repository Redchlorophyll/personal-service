package services

func NewService(config LinkyServiceConfig) LinkyServiceProvider {
	return &LinkyService{
		LinkyRepository:             config.LinkyRepository,
		ContentIdentifierRepository: config.ContentIdentifierRepository,
		Env:                         config.Env,
		AccountUrlSlugRepository:    config.AccountUrlSlugRepository,
		AccountService:              config.AccountService,
	}
}

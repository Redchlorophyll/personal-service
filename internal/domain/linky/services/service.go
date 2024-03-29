package services

func NewService(config LinkyServiceConfig) LinkyServiceProvider {
	return &LinkyService{
		LinkyRepository: config.LinkyRepository,
	}
}

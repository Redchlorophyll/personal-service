package httpservice

func NewHandler(config LinkyHandlerConfig) LinkyHandlerProvider {
	return &LinkyHandler{
		LinkyService: config.LinkyService,
	}
}

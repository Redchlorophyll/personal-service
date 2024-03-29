package main

import (
	env "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	PersonalServiceConfig "github.com/Redchlorophyll/personal-service/internal/config/personal_service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := env.GetEnvironmentVariables()

	app := fiber.New()
	serv := PersonalServiceConfig.GetService()
	httpService := PersonalServiceConfig.InitializeService(serv)

	httpService.LinkyHandler.SetRoute(app)

	app.Listen(config.Ports.PersonalServiceApi)
}

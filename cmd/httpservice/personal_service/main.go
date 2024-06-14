package main

import (
	env "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	PersonalServiceConfig "github.com/Redchlorophyll/personal-service/internal/config/personal_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config := env.GetEnvironmentVariables()

	app := fiber.New()
	serv := PersonalServiceConfig.GetService()
	httpService := PersonalServiceConfig.InitializeService(serv)

	httpService.LinkyHandler.SetRoute(app)
	httpService.AccountHandler.SetRoute(app)

	log.Info("starting serve on ", config.Ports.PersonalServiceApi, ". Env: ", config.Env)
	log.Fatal(app.Listen(":" + config.Ports.PersonalServiceApi))
}

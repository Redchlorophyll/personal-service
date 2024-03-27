package main

import (
	"fmt"

	internal "github.com/Redchlorophyll/personal-service/internal/config"
	PersonalServiceConfig "github.com/Redchlorophyll/personal-service/internal/config/personal_service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	serv := internal.GetService()
	httpService := PersonalServiceConfig.InitializeService(serv)

	fmt.Println(httpService)

	app.Listen(":4040")
}

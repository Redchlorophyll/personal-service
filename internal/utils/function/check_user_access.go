package function

import (
	env "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	utilsConstant "github.com/Redchlorophyll/personal-service/internal/utils/constant"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CheckUserAccess() func(*fiber.Ctx) error {
	return func(fiberContext *fiber.Ctx) error {
		token := fiberContext.Cookies("token")
		tokenFromEnv := env.GetEnvironmentVariables().SecretToken
		if token != tokenFromEnv {
			log.Error("[handler][CheckUserAccess] token given from Cookies are not match with secret token. ", token, tokenFromEnv)

			return fiberContext.Status(fiber.StatusForbidden).JSON(utilsResponse.GeneralResponse{
				StatusCode: fiber.StatusForbidden,
				Message:    utilsConstant.ERROR_MESSAGE[fiber.StatusForbidden].Error(),
			})
		}

		return fiberContext.Next()
	}
}

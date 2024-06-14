package services

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

func (service *AccountService) LogoutAccount(context context.Context, request string) error {
	err := service.AccountRepository.RevokeSessionToken(context, request)
	if err != nil {
		log.Error("[account][service][LogoutAccount] error when execute RevokeSessionToken(). ", err, context, request)

		return err
	}

	return nil
}

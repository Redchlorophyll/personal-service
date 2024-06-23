package services

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func (service *AccountService) VerifyTokenExpiration(context context.Context, sessionTokenExpiredAt *time.Time) error {
	if sessionTokenExpiredAt.Before(time.Now()) {
		log.Error("[account][service][VerifyTokenExpiration] error when execute service.VerifyTokenExpiration(). ", context, sessionTokenExpiredAt)
		return errors.New("unathorized access. session is expired")
	}

	return nil
}

package services

import (
	"context"
	"errors"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (service *AccountService) Login(context context.Context, req request.LoginRequest) error {
	user, err := service.AccountRepository.GetAccountByEmail(context, req.Email)
	if err != nil {
		log.Error("[account][service][Login] error when execute GetAccountByEmail(). ", err, context, req)

		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error("[account][service][Login] error when execute CompareHashAndPassword(). ", err, context, user.Password, req)

		return errors.New("[ERROR]: not found")
	}

	sessionToken := uuid.New()
	refreshToken := uuid.New()

	err = service.AccountRepository.SetSessionToken(context, request.SetSessionTokenRequest{
		SessionToken:    sessionToken.String(),
		RefreshToken:    refreshToken.String(),
		EmailIdentifier: req.Email,
	})
	if err != nil {
		log.Error("[account][service][Login] error when execute SetSessionToken(). ", err, context, req)

		return err
	}

	return nil
}

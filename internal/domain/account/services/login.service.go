package services

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (service *AccountService) Login(context context.Context, req request.LoginRequest) error {
	user, err := service.AccountRepository.GetAccountByEmail(context, req.Email)
	if err != nil {
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil
	}

	sessionToken := uuid.New()
	refreshToken := uuid.New()

	err = service.AccountRepository.SetSessionToken(context, request.SetSessionTokenRequest{
		SessionToken: sessionToken.String(),
		RefreshToken: refreshToken.String(),
	})

	return nil
}

package services

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/gofiber/fiber/v2/log"
)

func (service *AccountService) GetProfile(context context.Context, req string) (request.ProfileData, error) {
	account, err := service.AccountRepository.GetAccountBySessionToken(context, req)
	if err != nil {
		log.Error("[account][service][LogoutAccount] error when execute RevokeSessionToken(). ", err, context, req)

		return request.ProfileData{}, err
	}

	// check if session already expired.
	err = service.VerifyTokenExpiration(context, account.SessionTokenExpiredAt)

	if err != nil {
		log.Error("[account][service][LogoutAccount] error when execute service.VerifyTokenExpiration(). ", err, context, req)

		return request.ProfileData{}, err
	}

	return request.ProfileData{
		Username:    account.Username,
		Email:       account.Email,
		FullName:    account.FullName,
		Description: account.Description,
	}, nil
}

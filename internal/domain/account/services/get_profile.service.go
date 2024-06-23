package services

import (
	"context"
	"errors"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/constant"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (service *AccountService) GetProfile(context context.Context, req request.GetProfileServiceRequest) (request.ProfileData, error) {
	var (
		account response.GetAccountRespositoryResponse
		err     error
	)

	if req.Id != nil {
		account, err = service.AccountRepository.GetAccountsById(context, *req.Id)
		if err != nil {
			log.Error("[account][service][GetProfile] error when execute GetAccountBySessionToken(). ", err, context, req)

			return request.ProfileData{}, err
		}
	} else {
		account, err = service.AccountRepository.GetAccountBySessionToken(context, req.SessionToken)
		if err != nil {
			log.Error("[account][service][GetProfile] error when execute GetAccountBySessionToken(). ", err, context, req)

			return request.ProfileData{}, err
		}

		// check if session already expired.
		err = service.VerifyTokenExpiration(context, account.SessionTokenExpiredAt)
		if err != nil {
			log.Error("[account][service][GetProfile] error when execute service.VerifyTokenExpiration(). ", err, context, req)

			return request.ProfileData{}, errors.New("[ERROR]: unauthorized access")
		}
	}

	socialMedias, err := service.SocialMediaRepository.GetSocialMedias(context, request.GetSocialMediaRequest{
		AccountId: account.Id,
	})
	if err != nil {
		log.Error("[account][service][GetProfile] error when execute service.GetSocialMedias(). ", err, context, req, account)

		return request.ProfileData{}, err
	}

	MappedSocialMedias := make(map[constant.SocialMedia]string)
	for _, value := range socialMedias {
		MappedSocialMedias[value.SocialMedia] = value.Url
	}

	return request.ProfileData{
		Username:    account.Username,
		Email:       account.Email,
		FullName:    account.FullName,
		Description: account.Description,
		SocialMedia: MappedSocialMedias,
	}, nil
}

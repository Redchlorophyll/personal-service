package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Redchlorophyll/personal-service/internal/config/firebase"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	accountRequest "github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/utils/function"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func (service *AccountService) CreateAccount(context context.Context, request request.CreateAccountRequest) error {
	duplicateAccount, err := service.AccountRepository.GetAccountsByEmailOrUsername(context, accountRequest.GetAccountByEmailOrUsernameRequest{
		Email:    request.Email,
		Username: request.UserName,
	})
	if err != nil {
		log.Error("[account][service][CreateAccount] error when execute GetAccountsByEmailOrUsername(). ", err, context, request)
		return err
	}

	if len(duplicateAccount) != 0 {
		log.Error("[account][service][CreateAccount] error, duplicate account. ", err, context, request, duplicateAccount)
		return errors.New("[ERROR]: duplicate key")
	}

	err = function.ValidateFile(context, function.ValidateFileRequest{
		File:            request.ProfileImg,
		AllowedFileType: []string{"img/jpg", "img/png"},
		AllowedFileSize: 100,
	})
	if err != nil {
		log.Error("[account][service][CreateAccount] error when execute ValidateFile(). ", err, context, request)
		return err
	}

	uploadedImgUrls := make([]string, 1)
	var imgUrls string
	if request.ProfileImg != nil {
		folderPath := "profile-img/"
		uploadedImgUrls, err = service.FirebaseService.UploadFile(context, firebase.UploadFileToFirebaseRequest{
			FirebaseProject: "account_service",
			BucketName:      service.Env.Firebase["account_service"].StorageBucket,
			Files: []firebase.FilesRequest{
				firebase.FilesRequest{
					File:       request.ProfileImg,
					FolderPath: &folderPath,
				},
			},
		})
		if err != nil {
			log.Error("[account][service][CreateAccount] error when execute service.FirebaseService.UploadFile(). ", err, context, request)

			return err
		}

		imgUrls = uploadedImgUrls[0]
	} else {
		imgUrls = function.RandomizeProfileAvatar()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("[account][service][CreateAccount] error when execute GenerateFromPassword(). ", err, context, request)

		return err
	}

	createdAccount, err := service.AccountRepository.CreateAccount(context, accountRequest.CreateAccountRepositoryRequest{
		ProfileImgUrl: imgUrls,
		UserName:      request.UserName,
		FullName:      request.FullName,
		Description:   request.Description,
		Email:         request.Email,
		Password:      string(hashedPassword),
	})

	lowerCaseName := strings.ToLower(request.FullName)
	randString := function.RandomizeString(6)
	urlSlug := strings.ReplaceAll(lowerCaseName, " ", "-")

	err = service.AccountUrlSlugRepository.CreateAccountUrlSlug(context, accountRequest.CreateAccountUrlSlugRepositoryRequest{
		AccountId: int(createdAccount.Id),
		UrlSlug:   strings.Join([]string{urlSlug, randString}, "-"),
	})
	if err != nil {
		log.Error("[account][service][CreateAccount] error when execute service.AccountRepository.CreateAccountUrlSlug(). ", err, context, request, createdAccount)

		return errors.New("[ERROR]: partial error")
	}

	if err != nil {
		log.Error("[account][service][CreateAccount] error when execute service.AccountRepository.CreateAccount(). ", err, context, request)

		return err
	}

	return nil
}

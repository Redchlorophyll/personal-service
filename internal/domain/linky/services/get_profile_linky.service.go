package services

import (
	"context"

	requestAccount "github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	requestModel "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	utilsRequest "github.com/Redchlorophyll/personal-service/internal/utils/model/request"
	"github.com/gofiber/fiber/v2/log"
)

func (service *LinkyService) GetProfileLinky(context context.Context, request requestModel.GetProfileLinkyServiceRequest) (response.GetProfileLinkyServiceResponse, error) {
	var result response.GetProfileLinkyServiceResponse

	getAccountUrlSlugRequest := requestAccount.GetAccountUrlSlugRepositoryRequest{
		UrlSlug: request.UrlSlug,
	}
	urlSlug, err := service.AccountUrlSlugRepository.GetAccountUrlSlug(context, getAccountUrlSlugRequest)
	if err != nil {
		log.Error("[linky][services][GetProfileLinky] error when GetAccountUrlSlug(). ", err, request, getAccountUrlSlugRequest)

		return result, err
	}

	getProfileRequest := requestAccount.GetProfileServiceRequest{
		SessionToken: "",
		Id:           &urlSlug.AccountId,
	}
	account, err := service.AccountService.GetProfile(context, getProfileRequest)
	if err != nil {
		log.Error("[linky][services][GetProfileLinky] error when GetProfile(). ", err, request, getProfileRequest)
		return result, err
	}

	getIdentifierBySlugIdRequest := requestModel.GetIdentifierBySlugIdRepositoryRequest{
		UrlSlugId: int64(urlSlug.Id),
	}
	identifiers, err := service.ContentIdentifierRepository.GetIdentifierBySlugId(context, getIdentifierBySlugIdRequest)
	if err != nil {
		log.Error("[linky][services][GetProfileLinky] error when GetIdentifierBySlugId(). ", err, request, getIdentifierBySlugIdRequest)
		return result, err
	}

	linkyItems := []response.GetCleanLinkyResponse{}
	for _, identifier := range identifiers {
		var item response.GetCleanLinkyResponse
		getLinkyRequest := requestModel.GetLinkyRequestQuery{
			PaginationRequestQuery: utilsRequest.PaginationRequestQuery{
				// hardcoded for now.
				Page:    1,
				PerPage: 10,
			},
			Identifier: identifier.Identifier,
		}
		linky, err := service.GetLinky(context, getLinkyRequest)
		if err != nil {
			log.Error("[linky][services][GetProfileLinky] error when GetLinky(). ", err, request, getLinkyRequest)
			return result, err
		}

		item = response.GetCleanLinkyResponse{
			Metadata:   linky.Metadata,
			Data:       linky.Data,
			Pagination: linky.Pagination,
		}

		linkyItems = append(linkyItems, item)
	}

	result = response.GetProfileLinkyServiceResponse{
		ProfileInfo: account,
		Linky:       linkyItems,
	}

	return result, nil
}

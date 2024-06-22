package services

import (
	"context"

	requestModel "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
)

func (service *LinkyService) GetProfileLinky(context context.Context, request requestModel.GetLinkyRequestQuery) error {
	// PSUEDOCODE
	// 1. verify account by url slug
	// 2. if account exists, get the id
	// 3. get all identifier
	// 4. from that identifier id, loop to get linky.

	return nil
}

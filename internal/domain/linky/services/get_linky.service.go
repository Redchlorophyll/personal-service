package services

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
)

func (service *LinkyService) GetLinky(context context.Context, request request.GetLinkyRequestQuery) (response.GetLinkyResponse, error) {
	return response.GetLinkyResponse{}, nil
}

package social_media_table

import (
	"context"
	"database/sql"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
)

type SocialMediaTableRepositoryConfig struct {
	Db *sql.DB
}

type SocialmediaTableRepository struct {
	Db *sql.DB
}

type SocialMediaTableRepositoryProvider interface {
	GetSocialMedias(context context.Context, request request.GetSocialMediaRequest) ([]response.GetSocialMediaRepositoryResponse, error)
}

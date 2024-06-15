package response

import "github.com/Redchlorophyll/personal-service/internal/domain/account/model/constant"

type GetSocialMediaRepositoryResponse struct {
	SocialMedia constant.SocialMedia `db:"social_media"`
	Url         string               `db:"url"`
}

package request

import "github.com/Redchlorophyll/personal-service/internal/domain/account/model/constant"

type ProfileData struct {
	Username    string                          `json:"username"`
	Email       string                          `json:"email"`
	FullName    string                          `json:"full_name"`
	Description *string                         `json:"description"`
	SocialMedia map[constant.SocialMedia]string `json:"social_media"`
}

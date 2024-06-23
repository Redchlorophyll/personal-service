package response

import (
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
)

type GetProfileLinkyServiceResponse struct {
	ProfileInfo request.ProfileData     `json:"profile_info"`
	Linky       []GetCleanLinkyResponse `json:"linky"`
}

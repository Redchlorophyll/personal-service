package response

import (
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
)

type GetProfileResponse struct {
	utilsResponse.GeneralResponse

	Data *request.ProfileData `json:"data"`
}

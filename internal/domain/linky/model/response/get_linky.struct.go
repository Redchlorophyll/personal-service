package response

import (
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
)

type GetLinkyResponse struct {
	utilsResponse.GeneralResponse

	Metadata   *GetLinkyIdentifierResponse       `json:"metadata"`
	Data       []LinkyItem                       `json:"data"`
	Pagination *utilsResponse.PaginationResponse `json:"pagination"`
}

type LinkyItem struct {
	LinkyId     int     `db:"id" json:"linky_id"`
	Img         *string `db:"img_url" json:"img"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	UrlAnchor   string  `db:"url_anchor" json:"url_anchor"`
}

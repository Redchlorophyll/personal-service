package response

import (
	utilsResponse "github.com/Redchlorophyll/personal-service/internal/utils/model/response"
)

type GetLinkyResponse struct {
	Data       LinkyItem                        `json:"data"`
	Pagination utilsResponse.PaginationResponse `json:"pagination"`
}

type LinkyItem struct {
	LinkyId     int    `json:"linky_id"`
	Img         string `json:"img"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UrlAnchor   string `json:"url_anchor"`
}

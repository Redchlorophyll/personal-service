package response

type PaginationResponse struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	FromItem  int `json:"from_item"`
	ToItem    int `json:"to_item"`
	TotalItem int `json:"total_item"`
}

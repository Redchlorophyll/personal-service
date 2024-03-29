package request

type PaginationRequestQuery struct {
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}

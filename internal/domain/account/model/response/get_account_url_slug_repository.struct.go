package response

type GetAccountUrlSlugRepositoryResponse struct {
	Id        int    `db:"id"`
	UrlSlug   string `db:"url_slug"`
	AccountId int64  `db:"account_id"`
}

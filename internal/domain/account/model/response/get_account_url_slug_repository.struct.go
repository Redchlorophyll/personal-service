package response

type GetAccountUrlSlugRepositoryResponse struct {
	Id        int    `db:"id"`
	UrlSlug   string `db:"url_slug"`
	AccountId string `db:"account_id"`
}

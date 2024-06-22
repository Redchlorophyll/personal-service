package response

type GetIdentifierBySlugIdRepositoryResponse struct {
	Id         int64  `db:"id"`
	Identifier string `db:"identifier"`
	Title      string `db:"title"`
}

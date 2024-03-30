package response

type GetLinkyIdentifierResponse struct {
	IdentifierId int    `db:"id" json:"identifier_id"`
	Identifier   string `db:"identifier" json:"identifier"`
	Title        string `db:"title" json:"title"`
}

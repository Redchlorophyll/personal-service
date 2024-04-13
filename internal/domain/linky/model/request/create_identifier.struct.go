package request

type CreateIdentifierRequest struct {
	Identifier string `json:"identifier"`
	Title      string `json:"title"`
}

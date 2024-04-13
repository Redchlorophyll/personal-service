package request

type CreateLinkyRequest struct {
	IdentiferId int    `json:"identifier_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UrlAnchor   string `json:"url_anchhor" db:"url_anchor"`
	ImageUrl    string `json:"img_url" db:"img_url"`
}

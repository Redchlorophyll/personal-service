package request

import "errors"

type CreateLinkyRequest struct {
	IdentiferId int    `json:"identifier_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UrlAnchor   string `json:"url_anchor" db:"url_anchor"`
	ImageUrl    string `json:"img_url" db:"img_url"`
}

func (request *CreateLinkyRequest) Validator() error {
	if request.IdentiferId == 0 {
		return errors.New("invalid IdentiferId")
	}

	if request.Title == "" {
		return errors.New("invalid Title")
	}

	if request.Description == "" {
		return errors.New("invalid Description")
	}

	if request.UrlAnchor == "" {
		return errors.New("invalid UrlAnchor")
	}

	return nil
}

package linky_table

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type UpdateLinkyRequest struct {
	IdentiferId int    `json:"identifier_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UrlAnchor   string `json:"url_anchhor" db:"url_anchor"`
	ImageUrl    string `json:"img_url" db:"img_url"`
	Id          int
}

func (repository *LinkyTableRepository) UpdateLinky(context context.Context, request UpdateLinkyRequest) error {
	query := `
	UPDATE link
	SET 
		img_url = $1, 
		title = $2, 
		description = $3, 
		url_anchor = $4
	WHERE 
		id = $5
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request.ImageUrl,
		request.Title,
		request.Description,
		request.UrlAnchor,
		request.Id,
	)
	if err != nil {
		log.Error("[repository][UpdateLinky] error when ExecContext(). ", err, request)

		return err
	}

	return nil
}

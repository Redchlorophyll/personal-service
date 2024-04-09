package linky_table

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type CreateLinkyRequest struct {
	IdentiferId int    `json:"identifier_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UrlAnchor   string `json:"url_anchhor" db:"url_anchor"`
	ImageUrl    string `json:"img_url" db:"img_url"`
}

func (repository LinkyTableRepository) CreateLinky(context context.Context, request CreateLinkyRequest) (*int, error) {
	var result int

	query := `
		INSERT INTO 
			link 
			(img_url, title, description, url_anchor)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`

	statement, err := repository.Db.PrepareContext(
		context,
		query,
	)
	if err != nil {
		log.Error("[repository][CreateLinky] error when PrepareContext(). ", err, request)

		return &result, err
	}

	defer statement.Close()

	err = statement.QueryRow(
		request.ImageUrl,
		request.Title,
		request.Description,
		request.UrlAnchor,
	).Scan(&result)
	if err != nil {
		log.Error("[repository][CreateLinky] error when QueryRow().Scan(). ", err, request)
		return &result, err
	}

	return &result, nil
}

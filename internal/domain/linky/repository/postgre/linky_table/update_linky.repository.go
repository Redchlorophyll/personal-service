package linky_table

import (
	"context"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/gofiber/fiber/v2/log"
)

func (repository *LinkyTableRepository) UpdateLinky(context context.Context, request modelRequest.UpdateLinkyRequest) error {
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

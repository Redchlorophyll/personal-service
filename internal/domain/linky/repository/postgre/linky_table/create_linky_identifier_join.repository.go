package linky_table

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type CreateLinkyIdentifierRequest struct {
	ContentIdentiferId int
	LinkId             int
}

func (repository LinkyTableRepository) CreateLinkyIdentifier(context context.Context, request CreateLinkyIdentifierRequest) error {
	query := `
		INSERT INTO 
			link_identifier 
			(type_id, link_id)
		VALUES
			($1, $2)
		RETURNING id
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request.ContentIdentiferId,
		request.LinkId,
	)
	if err != nil {
		log.Error("[repository][CreateLinky] error when CreateLinkyIdentifier(). ", err, request)

		return err
	}

	return nil
}

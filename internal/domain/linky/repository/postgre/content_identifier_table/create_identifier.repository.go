package content_identifier_table

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type CreateIdentifierRequest struct {
	Identifier string
	Title      string
}

func (repository *ContentIdentifierTableRepository) CreateIdentifier(context context.Context, request CreateIdentifierRequest) error {
	query := `
		INSERT INTO 
			content_identifier 
			(identifier, title)
		VALUES
			($1, $2)
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request.Identifier,
		request.Title,
	)
	if err != nil {
		log.Error("[repository][CreateIdentifier] error when ExecContext(). ", err, request)

		return err
	}

	return nil
}

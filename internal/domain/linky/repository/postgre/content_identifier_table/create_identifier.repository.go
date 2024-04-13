package content_identifier_table

import (
	"context"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/gofiber/fiber/v2/log"
)

func (repository *ContentIdentifierTableRepository) CreateIdentifier(context context.Context, request modelRequest.CreateIdentifierRequest) error {
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

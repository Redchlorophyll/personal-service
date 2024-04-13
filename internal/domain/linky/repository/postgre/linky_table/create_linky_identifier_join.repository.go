package linky_table

import (
	"context"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/gofiber/fiber/v2/log"
)

func (repository LinkyTableRepository) CreateLinkyIdentifier(context context.Context, request modelRequest.CreateLinkyIdentifierRequest) error {
	query := `
		INSERT INTO 
			link_identifier 
			(type_id, link_id)
		VALUES
			($1, $2)
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request.ContentIdentiferId,
		request.LinkId,
	)
	if err != nil {
		log.Error("[repository][CreateLinkyIdentifier] error when ExecContext(). ", err, request)

		return err
	}

	return nil
}

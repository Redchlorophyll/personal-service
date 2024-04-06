package linky_table

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository LinkyTableRepository) GetLinkyIdentifier(context context.Context, request string) (response.GetLinkyIdentifierResponse, error) {
	var result response.GetLinkyIdentifierResponse

	if request == "" {
		return result, nil
	}

	query := `
		SELECT
			id,
			identifier,
			title
		FROM
			content_identifier
		WHERE
			identifier = $1

	`

	rows, err := repository.Db.QueryContext(context, query, request)
	if err != nil {
		log.Error("[repository][GetLinkyIdentifier] error when QueryContext. ", err, request)

		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Next()

		err = rows.Scan(&result.IdentifierId, &result.Identifier, &result.Title)
		if err != nil {
			log.Error("[repository][GetLinkyIdentifier] error when rows.Scan(). ", err, request)

			return result, err
		}
	}

	return result, nil
}

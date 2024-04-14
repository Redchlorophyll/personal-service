package linky_table

import (
	"context"

	modelResponse "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository LinkyTableRepository) GetLinkyIdentifierById(context context.Context, request *int) (modelResponse.GetLinkyIdentifierResponse, error) {
	var result modelResponse.GetLinkyIdentifierResponse

	if request == nil {
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
			id = $1

	`

	rows, err := repository.Db.QueryContext(context, query, request)
	if err != nil {
		log.Error("[repository][GetLinkyIdentifierById] error when QueryContext(). ", err, request)

		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.IdentifierId, &result.Identifier, &result.Title)
		if err != nil {
			log.Error("[repository][GetLinkyIdentifierById] error when Scan(). ", err, request)

			return result, err
		}
	}

	return result, nil
}

package content_identifier_table

import (
	"context"

	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository *ContentIdentifierTableRepository) GetIdentifierBySlugId(context context.Context, request modelRequest.GetIdentifierBySlugIdRepositoryRequest) ([]response.GetIdentifierBySlugIdRepositoryResponse, error) {
	results := []response.GetIdentifierBySlugIdRepositoryResponse{}

	query := `
		SELECT
			id,
			identifier,
			title
		FROM
			content_identifier
		WHERE
			account_url_slug_id = $1;
	`

	rows, err := repository.Db.QueryContext(
		context,
		query,
		request.UrlSlugId,
	)
	if err != nil {
		log.Error("[linky][repository][GetIdentifierBySlugId] error when QueryContext(). ", err, request)

		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		var item response.GetIdentifierBySlugIdRepositoryResponse

		err = rows.Scan(&item.Id, &item.Identifier, &item.Title)
		if err != nil {
			log.Panic("[linky][repository][GetIdentifierBySlugId] error when Scan(). ", err, request)
		}

		results = append(results, item)
	}

	return results, nil
}

package linky_table

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	utilsFunction "github.com/Redchlorophyll/personal-service/internal/utils/function"
	"github.com/gofiber/fiber/v2/log"
)

func (repository LinkyTableRepository) GetLinky(context context.Context, request request.GetLinkyRequestQuery) ([]response.LinkyItem, error) {
	results := []response.LinkyItem{}
	args := []interface{}{}

	offset := utilsFunction.GetPaginationOffset(request.Page)

	query := `
		SELECT
			l.id,
			l.img_url,
			l.title,
			l.description,
			l.url_anchor
		FROM
			link AS l
		JOIN
			link_identifier AS li ON l.id = li.link_id
		JOIN
			content_identifier AS ci ON li.type_id = ci.id
		WHERE
			l.deleted_at IS NULL
	`

	if request.Identifier != "" {
		query += `
		AND
			ci.identifier = $1
		ORDER BY
			id DESC
		LIMIT
			$2
		OFFSET
			$3
		`
		args = append(args, request.Identifier, request.PerPage, offset)
	} else {
		query += `
		ORDER BY
			id DESC
		LIMIT
			$1
		OFFSET
			$2
		`

		args = append(args, request.PerPage, offset)
	}

	rows, err := repository.Db.QueryContext(context, query, args...)
	if err != nil {
		log.Error("[repository][GetLinky] error when QueryContext. ", err, request)

		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		var item response.LinkyItem

		err := rows.Scan(
			&item.LinkyId,
			&item.Img,
			&item.Title,
			&item.Description,
			&item.UrlAnchor,
		)
		if err != nil {
			log.Panic("[repository][GetLinky] error when rows.Scan(). ", err, request)
		}

		results = append(results, item)
	}

	return results, nil
}

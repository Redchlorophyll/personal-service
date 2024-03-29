package linky_table

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
	utilsFunction "github.com/Redchlorophyll/personal-service/internal/utils/function"
	utilsRequest "github.com/Redchlorophyll/personal-service/internal/utils/model/request"
)

type LinkyRepositoryResponse struct {
	LinkyId     int    `db:"id"`
	Img         string `db:"img_url"`
	Title       string `db:"title"`
	Description string `db:"description"`
	UrlAnchor   string `db:"url_anchor"`
}

func (repository LinkyTableRepository) GetLinky(context context.Context, pagination utilsRequest.PaginationRequestQuery) ([]response.LinkyItem, error) {
	results := []response.LinkyItem{}

	offset := utilsFunction.GetPaginationOffset(pagination.Page)

	query := `
		SELECT
			id,
			img,
			title,
			description,
			url_anchor
		FROM
			link
		WHERE
			deleted_at IS NULL
		ORDER BY
			id DESC
		LIMIT
			$1
		OFFSET
			$2
	`

	rows, err := repository.Db.QueryContext(context, query, pagination.PerPage, offset)
	if err != nil {
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
			panic(err)
		}

		results = append(results, item)
	}

	return results, nil
}

func (repository LinkyTableRepository) GetTotalLinkyItem(context context.Context) (int, error) {

	query := `
		SELECT
			COUNT(id)
		FROM
			link
		WHERE
			deleted_at IS NULL
	`

	rows, err := repository.Db.QueryContext(context, query)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var totalItems int
	rows.Scan(&totalItems)

	return totalItems, nil
}

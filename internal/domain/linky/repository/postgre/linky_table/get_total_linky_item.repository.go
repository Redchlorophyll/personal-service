package linky_table

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

func (repository LinkyTableRepository) GetTotalLinkyItem(context context.Context, request string) (int, error) {
	args := []interface{}{}

	query := `
		SELECT
			COUNT(l.id)
		FROM
			link AS l
		JOIN
			link_identifier AS li ON l.id = li.link_id
		JOIN
			content_identifier AS ci ON li.type_id = ci.id
		WHERE
			l.deleted_at IS NULL
	`

	if request != "" {
		query += `
		AND
			ci.identifier = $1
		`
		args = append(args, request)
	}

	rows, err := repository.Db.QueryContext(context, query, args...)
	if err != nil {
		log.Error("[repository][GetTotalLinkyItem] error when QueryContext(). ", err)
		return 0, err
	}

	defer rows.Close()

	var totalItems int

	for rows.Next() {
		err = rows.Scan(&totalItems)
		if err != nil {
			log.Error("[repository][GetTotalLinkyItem] error when Scan(). ", err)

			return totalItems, err
		}
	}

	return totalItems, nil
}

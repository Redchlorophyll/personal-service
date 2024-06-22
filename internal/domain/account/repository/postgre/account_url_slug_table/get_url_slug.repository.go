package account_url_slug_table

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository *AccountUrlSlugTableRepository) GetAccountUrlSlug(context context.Context, request request.GetAccountUrlSlugRepositoryRequest) (response.GetAccountUrlSlugRepositoryResponse, error) {
	var result response.GetAccountUrlSlugRepositoryResponse
	query := `
		SELECT
			id,
			url_slug,
			account_id
		FROM
			account_url_slug_table
		WHERE
			url_slug = $1
	`

	rows := repository.Db.QueryRowContext(context, query, request.UrlSlug)

	err := rows.Scan(
		&result.Id,
		&result.UrlSlug,
		&result.AccountId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
			log.Warn("[account][repository][GetUrlSlug] no rows found for url_slug: ", context, err, request)
			return result, errors.New("[ERROR]: not found")
		}

		log.Error("[account][repository][GetUrlSlug] error when Scan(). ", context, err, request)
		return result, err
	}

	return result, nil
}

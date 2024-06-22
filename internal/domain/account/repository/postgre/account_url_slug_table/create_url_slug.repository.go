package account_url_slug_table

import (
	"context"
	"errors"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/gofiber/fiber/v2/log"
	"github.com/lib/pq"
)

func (repository *AccountUrlSlugTableRepository) CreateAccountUrlSlug(context context.Context, request request.CreateAccountUrlSlugRepositoryRequest) error {
	query := `
		INSERT INTO
			account_url_slug_table
			(
				url_slug,
				account_id
			)
		VALUES
			($1, $2)
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request.UrlSlug,
		request.AccountId,
	)
	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		log.Error("[account][repository][CreateAccountUrlSlug] error duplicate key. ", err, request)
		return errors.New("[ERROR]: duplicate key")
	}

	log.Error("[account][repository][CreateAccountUrlSlug] error when ExecContext(). ", err, request)
	return err
}

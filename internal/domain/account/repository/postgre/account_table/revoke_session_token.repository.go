package account_table

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

func (repository AccountTableRepository) RevokeSessionToken(context context.Context, request string) error {
	query := `
		UPDATE 
			account_table
		SET 
			last_logged_out = NOW(),
			session_token = null,
			session_token_expired_at = null,
			refresh_token = null,
			refresh_token_expired_at = null
		WHERE 
			session_token = $1
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request,
	)

	if err != nil {
		log.Error("[account][repository][RevokeSessionToken] error when ExecContext(). ", err, request)

		return err
	}

	return nil
}

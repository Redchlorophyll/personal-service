package account_table

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/gofiber/fiber/v2/log"
)

func (repository AccountTableRepository) SetSessionToken(context context.Context, request request.SetSessionTokenRequest) error {
	query := `
		UPDATE 
			account_table
		SET 
			last_logged_in = NOW(),
			session_token = $1,
			session_token_expired_at = NOW() + INTERVAL '2 days',
			refresh_token = $2,
			refresh_token_expired_at = NOW() + INTERVAL '5 days'
		WHERE 
			email = $3
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request.SessionToken,
		request.RefreshToken,
		request.EmailIdentifier,
	)

	if err != nil {
		log.Error("[account][repository][SetSessionToken] error when ExecContext(). ", err, request)

		return err
	}

	return nil
}

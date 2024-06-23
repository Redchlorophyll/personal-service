package account_table

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository AccountTableRepository) GetAccountBySessionToken(context context.Context, request string) (response.GetAccountRespositoryResponse, error) {
	var result response.GetAccountRespositoryResponse

	query := `
		SELECT
			id,
			session_token_expired_at,
			session_token,
			user_name,
			email,
			full_name,
			description,
			password
		FROM
			account_table
		WHERE
			session_token = $1;
	`

	rows := repository.Db.QueryRowContext(context, query, request)

	err := rows.Scan(
		&result.Id,
		&result.SessionTokenExpiredAt,
		&result.SessionToken,
		&result.Username,
		&result.Email,
		&result.FullName,
		&result.Description,
		&result.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
			log.Warn("[account][repository][GetAccountBySessionToken] no rows found for email: ", context, err, request)
			return result, errors.New("[ERROR]: not found")
		}

		log.Error("[account][repository][GetAccountByEmail] error when Scan(). ", context, err, request)
		return result, err
	}

	return result, nil
}

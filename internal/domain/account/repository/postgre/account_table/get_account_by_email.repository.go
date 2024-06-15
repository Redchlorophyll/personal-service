package account_table

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository AccountTableRepository) GetAccountByEmail(context context.Context, email string) (response.GetAccountRespositoryResponse, error) {
	var result response.GetAccountRespositoryResponse

	query := `
		SELECT
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
			email = $1;
	`

	rows := repository.Db.QueryRowContext(context, query, email)

	err := rows.Scan(
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
			log.Warn("[account][repository][GetAccountByEmail] no rows found for email: ", email)
			return result, errors.New("[ERROR]: not found")
		}

		log.Error("[account][repository][GetAccountByEmail] error when Scan(). ", err, email)
		return result, err
	}

	return result, nil
}

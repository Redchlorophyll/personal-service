package account_table

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository AccountTableRepository) GetAccountBySessionToken(context context.Context, request string) (response.GetAccountRespositoryResponse, error) {
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
			session_token = $1;
	`

	rows, err := repository.Db.QueryContext(context, query, request)

	if err != nil {
		log.Error("[account][repository][GetAccount] error when QueryContext(). ", err, request)

		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&result.SessionTokenExpiredAt,
			&result.SessionToken,
			&result.Username,
			&result.Email,
			&result.FullName,
			&result.Description,
			&result.Password,
		)

		if err != nil {
			log.Error("[repository][GetLinkyIdentifier] error when Scan(). ", err, request)

			return result, err
		}
	}

	return result, nil
}

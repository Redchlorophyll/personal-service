package account_table

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository AccountTableRepository) GetAccountsById(context context.Context, requestId int64) (response.GetAccountRespositoryResponse, error) {
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
			id = $1
	`

	rows := repository.Db.QueryRowContext(context, query, requestId)

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
		log.Error("[account][repository][GetAccountsById] error when Scan(). ", err, requestId)

		return result, err
	}

	return result, nil
}

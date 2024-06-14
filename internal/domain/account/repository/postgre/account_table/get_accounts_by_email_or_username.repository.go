package account_table

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
)

func (repository AccountTableRepository) GetAccountsByEmailOrUsername(context context.Context, request request.GetAccountByEmailOrUsernameRequest) ([]response.GetAccountRespositoryResponse, error) {
	results := []response.GetAccountRespositoryResponse{}

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
			email = $1
		OR
			user_name = $2;
	`

	rows, err := repository.Db.QueryContext(context, query, request.Email, request.Username)

	if err != nil {
		log.Error("[account][repository][GetAccountByEmailOrUsername] error when QueryContext(). ", err, request)

		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		var result response.GetAccountRespositoryResponse

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
			log.Error("[account][repository][GetAccountByEmailOrUsername] error when Scan(). ", err, request)

			return results, err
		}

		results = append(results, result)
	}

	return results, nil
}

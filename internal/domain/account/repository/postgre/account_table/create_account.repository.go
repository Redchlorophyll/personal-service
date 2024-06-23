package account_table

import (
	"context"
	"errors"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
	"github.com/lib/pq"
)

func (repository AccountTableRepository) CreateAccount(context context.Context, request request.CreateAccountRepositoryRequest) (response.CreateAccountRepositoryResponse, error) {
	var result response.CreateAccountRepositoryResponse
	query := `
		INSERT INTO
			account_table
			(
				email, 
				profile_img_url, 
				user_name, 
				description, 
				full_name,
				password,
				created_at
			)
		VALUES
			($1, $2, $3, $4, $5, $6, NOW())
		RETURNING
			id
	`

	err := repository.Db.QueryRowContext(
		context,
		query,
		request.Email,
		request.ProfileImgUrl,
		request.UserName,
		request.Description,
		request.FullName,
		request.Password,
	).Scan(&result.Id)

	if err == nil {
		return result, nil
	}

	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		log.Error("[account][repository][CreateAccount] error duplicate key. ", err, request)
		return result, errors.New("[ERROR]: duplicate key")
	}

	log.Error("[account][repository][CreateAccount] error when ExecContext(). ", err, request)
	return result, err
}

package account_table

import (
	"context"
	"errors"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/gofiber/fiber/v2/log"
	"github.com/lib/pq"
)

func (repository AccountTableRepository) CreateAccount(context context.Context, request request.CreateAccountRepositoryRequest) error {
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
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request.Email,
		request.ProfileImgUrl,
		request.UserName,
		request.Description,
		request.FullName,
		request.Password,
	)

	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		log.Error("[account][repository][CreateAccount] error duplicate key. ", err, request)
		return errors.New("[ERROR]: duplicate key")
	}

	log.Error("[account][repository][CreateAccount] error when ExecContext(). ", err, request)
	return err
}

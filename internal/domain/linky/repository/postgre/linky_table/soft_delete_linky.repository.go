package linky_table

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

func (repository *LinkyTableRepository) SoftDeleteLinky(context context.Context, request int) error {
	query := `
	UPDATE link
	SET 
		deleted_at = NOW()
	WHERE 
		id = $1
	`

	_, err := repository.Db.ExecContext(
		context,
		query,
		request,
	)
	if err != nil {
		log.Error("[repository][SoftDeleteLinky] error when ExecContext(). ", err, request)

		return err
	}

	return nil
}

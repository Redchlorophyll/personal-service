package social_media_table

import (
	"context"

	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/account/model/response"
	"github.com/gofiber/fiber/v2/log"
)

// thinking possibility of joinging search account in get_account. then join social media table there.
func (repository *SocialmediaTableRepository) GetSocialMedias(context context.Context, request request.GetSocialMediaRequest) ([]response.GetSocialMediaRepositoryResponse, error) {
	var results []response.GetSocialMediaRepositoryResponse

	query := `
		SELECT
			social_media,
			url
		FROM
			social_media_table
		WHERE
			account_id = $1
	`

	rows, err := repository.Db.QueryContext(context, query, request.AccountId)
	if err != nil {
		log.Error("[repository][GetLinkyIdentifier] error when QueryContext. ", err, request)

		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		var item response.GetSocialMediaRepositoryResponse

		err = rows.Scan(&item.SocialMedia, &item.Url)
		if err != nil {
			log.Panic("[repository][GetLinky] error when Scan(). ", err, request)
		}

		results = append(results, item)
	}

	return results, nil
}

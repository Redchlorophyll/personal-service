package personal_service

import (
	"github.com/Redchlorophyll/personal-service/internal/config/database"
	env "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	contentRepositoryTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/content_identifier_table"
	linkyTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/linky_table"
	LinkyService "github.com/Redchlorophyll/personal-service/internal/domain/linky/services"
)

func GetService() *Service {
	env := env.GetEnvironmentVariables()
	db := database.InitializeDatabase(env)

	LinkyRepository := linkyTable.NewLinkyTableRepository(
		linkyTable.LinkyTableRepositoryConfig{Db: db["personal_service"]},
	)
	ContentIdentifierRepository := contentRepositoryTable.NewContentIdentifierTableRepository(
		contentRepositoryTable.ContentIdentifierTableRepositoryConfig{Db: db["personal_service"]},
	)

	return &Service{
		LinkyService: LinkyService.NewService(LinkyService.LinkyServiceConfig{
			LinkyRepository:             LinkyRepository,
			ContentIdentifierRepository: ContentIdentifierRepository,
		}),
	}
}

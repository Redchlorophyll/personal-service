package personal_service

import (
	"github.com/Redchlorophyll/personal-service/internal/config/database"
	env "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	AccountTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/account_table"
	AccountService "github.com/Redchlorophyll/personal-service/internal/domain/account/services"
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
	AccountRepository := AccountTable.NewAccountTableRepository(
		AccountTable.AccountTableRepositoryConfig{Db: db["personal_service"]},
	)
	ContentIdentifierRepository := contentRepositoryTable.NewContentIdentifierTableRepository(
		contentRepositoryTable.ContentIdentifierTableRepositoryConfig{Db: db["personal_service"]},
	)

	return &Service{
		LinkyService: LinkyService.NewService(LinkyService.LinkyServiceConfig{
			LinkyRepository:             LinkyRepository,
			ContentIdentifierRepository: ContentIdentifierRepository,
		}),
		AccountService: AccountService.NewAccountService(AccountService.AccountServiceConfig{
			AccountRepository: AccountRepository,
		}),
	}
}

package personal_service

import (
	"github.com/Redchlorophyll/personal-service/internal/config/database"
	env "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	"github.com/Redchlorophyll/personal-service/internal/config/firebase"
	AccountTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/account_table"
	AccountUrlSlugTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/account_url_slug_table"
	SocialMediaTable "github.com/Redchlorophyll/personal-service/internal/domain/account/repository/postgre/social_media_table"
	AccountService "github.com/Redchlorophyll/personal-service/internal/domain/account/services"
	contentRepositoryTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/content_identifier_table"
	linkyTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/linky_table"
	LinkyService "github.com/Redchlorophyll/personal-service/internal/domain/linky/services"
)

func GetService() *Service {
	env := env.GetEnvironmentVariables()
	db := database.InitializeDatabase(env)
	firebaseService := firebase.NewFirebase(firebase.FirebaseServiceConfig{
		Env: env,
	})

	// repositories
	linkyRepository := linkyTable.NewLinkyTableRepository(
		linkyTable.LinkyTableRepositoryConfig{
			Db: db["personal_service"],
		},
	)
	accountRepository := AccountTable.NewAccountTableRepository(
		AccountTable.AccountTableRepositoryConfig{
			Db: db["personal_service"],
		},
	)
	accountUrlSlugRepository := AccountUrlSlugTable.NewUrlSlugTable(
		AccountUrlSlugTable.AccountUrlSlugTableRepositoryConfig{
			Db: db["personal_service"],
		},
	)
	socialMediaRepository := SocialMediaTable.NewSocialMediaTableRepository(
		SocialMediaTable.SocialMediaTableRepositoryConfig{
			Db: db["personal_service"],
		},
	)
	contentIdentifierRepository := contentRepositoryTable.NewContentIdentifierTableRepository(
		contentRepositoryTable.ContentIdentifierTableRepositoryConfig{Db: db["personal_service"]},
	)

	// services
	AccountService := AccountService.NewAccountService(AccountService.AccountServiceConfig{
		AccountRepository:        accountRepository,
		FirebaseService:          firebaseService,
		SocialMediaRepository:    socialMediaRepository,
		AccountUrlSlugRepository: accountUrlSlugRepository,
		Env:                      env,
	})

	return &Service{
		LinkyService: LinkyService.NewService(LinkyService.LinkyServiceConfig{
			LinkyRepository:             linkyRepository,
			ContentIdentifierRepository: contentIdentifierRepository,
			AccountUrlSlugRepository:    accountUrlSlugRepository,
			Env:                         env,

			// services deps.
			AccountService: AccountService,
		}),
		AccountService: AccountService,
	}
}

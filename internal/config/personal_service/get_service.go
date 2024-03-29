package personal_service

import (
	"github.com/Redchlorophyll/personal-service/internal/config/database"
	env "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	linkyTable "github.com/Redchlorophyll/personal-service/internal/domain/linky/repository/postgre/linky_table"
	LinkyService "github.com/Redchlorophyll/personal-service/internal/domain/linky/services"
)

func GetService() *Service {
	env := env.GetEnvironmentVariables()
	db := database.InitializeDatabase(env)

	LinkyRepository := linkyTable.NewLinkyTableRepository(linkyTable.LinkyTableRepositoryConfig{Db: db["personal-service"]})

	return &Service{
		LinkyService: LinkyService.NewService(LinkyService.LinkyServiceConfig{
			LinkyRepository: LinkyRepository,
		}),
	}
}

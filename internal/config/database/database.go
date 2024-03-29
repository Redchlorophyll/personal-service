package database

import (
	"database/sql"
	"time"

	envVariable "github.com/Redchlorophyll/personal-service/internal/config/environment_variable"
	_ "github.com/lib/pq"
)

func InitializeDatabase(config envVariable.Config) map[string]*sql.DB {
	var res = make(map[string]*sql.DB)

	for key, variable := range config.Database {
		db, err := sql.Open(variable.Driver, variable.Master)

		if err != nil {
			panic(err)
		}

		db.SetMaxIdleConns(variable.MaxIdleConn)
		db.SetMaxOpenConns(variable.MaxOpenConn)
		db.SetConnMaxIdleTime(time.Duration(variable.MaxIdleTime.Minutes()))
		db.SetConnMaxLifetime(time.Duration(variable.MaxLifeTime.Minutes()))

		res[key] = db
	}

	return res
}

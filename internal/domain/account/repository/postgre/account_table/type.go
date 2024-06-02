package account_table

import "database/sql"

type AccountTableRepository struct {
	Db *sql.DB
}

type AccountTableRepositoryConfig struct {
	Db *sql.DB
}

type AccountTableRepositoryProvider interface{}

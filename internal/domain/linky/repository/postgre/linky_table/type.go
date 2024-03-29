package linky_table

import "database/sql"

type LinkyTableRepository struct {
	Db *sql.DB
}

type LinkyTableRepositoryConfig struct {
	Db *sql.DB
}

type LinkyTableRepositoryProvider interface {
	GetLinky() error
}

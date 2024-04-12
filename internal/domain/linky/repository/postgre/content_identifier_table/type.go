package content_identifier_table

import (
	"context"
	"database/sql"
)

type ContentIdentifierTableRepository struct {
	Db *sql.DB
}

type ContentIdentifierTableRepositoryConfig struct {
	Db *sql.DB
}

type ContentIdentifierTableRepositoryProvider interface {
	CreateIdentifier(context context.Context, request CreateIdentifierRequest) error
}

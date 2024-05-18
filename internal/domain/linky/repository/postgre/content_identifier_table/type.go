package content_identifier_table

import (
	"context"
	"database/sql"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
)

type ContentIdentifierTableRepository struct {
	Db *sql.DB
}

type ContentIdentifierTableRepositoryConfig struct {
	Db *sql.DB
}

type ContentIdentifierTableRepositoryProvider interface {
	CreateIdentifier(context context.Context, request request.CreateIdentifierRequest) error
}

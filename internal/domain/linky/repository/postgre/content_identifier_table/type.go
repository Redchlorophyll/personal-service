package content_identifier_table

import (
	"context"
	"database/sql"

	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	modelRequest "github.com/Redchlorophyll/personal-service/internal/domain/linky/model/request"
	"github.com/Redchlorophyll/personal-service/internal/domain/linky/model/response"
)

type ContentIdentifierTableRepository struct {
	Db *sql.DB
}

type ContentIdentifierTableRepositoryConfig struct {
	Db *sql.DB
}

type ContentIdentifierTableRepositoryProvider interface {
	CreateIdentifier(context context.Context, request request.CreateIdentifierRequest) error
	GetIdentifierBySlugId(context context.Context, request modelRequest.GetIdentifierBySlugIdRepositoryRequest) ([]response.GetIdentifierBySlugIdRepositoryResponse, error)
}

package content_identifier_table

func NewContentIdentifierTableRepository(config ContentIdentifierTableRepositoryConfig) ContentIdentifierTableRepositoryProvider {
	return &ContentIdentifierTableRepository{
		Db: config.Db,
	}
}

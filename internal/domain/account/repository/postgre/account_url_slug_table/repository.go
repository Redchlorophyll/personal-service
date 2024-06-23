package account_url_slug_table

func NewUrlSlugTable(config AccountUrlSlugTableRepositoryConfig) AccountUrlSlugTableRepositoryProvider {
	return &AccountUrlSlugTableRepository{
		Db: config.Db,
	}
}

package account_table

func NewAccountTableRepository(config AccountTableRepositoryConfig) AccountTableRepositoryProvider {
	return &AccountTableRepository{
		Db: config.Db,
	}
}

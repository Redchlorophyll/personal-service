package linky_table

func NewLinkyTableRepository(config LinkyTableRepositoryConfig) LinkyTableRepositoryProvider {
	return &LinkyTableRepository{
		Db: config.Db,
	}
}

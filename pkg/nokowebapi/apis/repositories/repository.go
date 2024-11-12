package repositories

type BaseRepositoryImpl struct{}

type BaseRepository struct{}

func NewBaseRepository() BaseRepositoryImpl {
	return BaseRepositoryImpl{}
}

package repositories

type BaseRepositoryImpl interface{}

type BaseRepository struct{}

func NewBaseRepository() BaseRepositoryImpl {
	return &BaseRepository{}
}

package repository

type  SongRepository interface {
}

type Repository struct {
	UserRepository
}

func NewRepository() *Repository {
	return &Repository{
		UserRepository: NewUserTarantool(),
		
	}
}
package service

type UserService interface {
}

type Service struct {
	UserService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repos.SongRepository),
	}
}
package service

type UserServiceImpl struct {
    repo       repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
    return &UserServiceImpl{
        repo:       repo,
    }
}
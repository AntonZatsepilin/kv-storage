package service

import (
	"github.com/AntonZatsepilin/kv-storage.git/internal/repository"
)

type KeyValueService interface {
}

type Service struct {
	KeyValueService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		KeyValueService: NewKeyValueService(repo.KeyValueRepository),
	}
}

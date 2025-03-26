package service

import (
	"github.com/AntonZatsepilin/kv-storage.git/internal/repository"
)

type KeyValueService interface {
	SetValue(key, value string) error
	GetValueByKey(key string) (string, error)
	UpdateValue(key, value string) error
	DeleteValue(key string) error 
}

type Service struct {
	KeyValueService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		KeyValueService: NewKeyValueService(repo.KeyValueRepository),
	}
}

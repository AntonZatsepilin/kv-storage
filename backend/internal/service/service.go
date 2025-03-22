package service

import (
	"errors"

	"github.com/AntonZatsepilin/kv-storage.git/internal/repository"
)

var (
	ErrKeyExists   = errors.New("key already exists")
	ErrKeyNotFound = errors.New("key not found")
)

type KVService interface {
	Create(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Update(key string, value interface{}) error
	Delete(key string) error
}

type Service struct {
	KVService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		KVService: NewKVServiceImpl(repo.KVRepository),
	}
}

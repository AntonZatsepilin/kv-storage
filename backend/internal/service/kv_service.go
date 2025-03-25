package service

import "github.com/AntonZatsepilin/kv-storage.git/internal/repository"

type KeyValueServiceImpl struct {
	repo repository.KeyValueRepository
}

func NewKeyValueService(repo repository.KeyValueRepository) *KeyValueServiceImpl {
	return &KeyValueServiceImpl{repo: repo}
}

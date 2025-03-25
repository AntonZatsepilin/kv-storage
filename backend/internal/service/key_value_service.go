package service

import "github.com/AntonZatsepilin/kv-storage.git/internal/repository"

type KeyValueServiceImpl struct {
	repo repository.KeyValueRepository
}

func NewKeyValueService(repo repository.KeyValueRepository) *KeyValueServiceImpl {
	return &KeyValueServiceImpl{repo: repo}
}

func (s *KeyValueServiceImpl) SetValue(key, value string) error {
	return s.repo.SetValue(key, value)
}

func (s *KeyValueServiceImpl) GetValueByKey(key string) (string, error) {
	return s.repo.GetValueByKey(key)
}

func (s *KeyValueServiceImpl) UpdateValue(key, value string) error {
	return s.repo.SetValue(key, value)
}
package service

import "github.com/AntonZatsepilin/kv-storage.git/internal/repository"

type KVServiceImpl struct {
	repo repository.KVRepository
}

func NewKVServiceImpl(repo repository.KVRepository) *KVServiceImpl {
	return &KVServiceImpl{repo: repo}
}

func (s *KVServiceImpl) Create(key string, value interface{}) error {
	err := s.repo.Create(key, value)
	if err != nil && err.Error() == "key already exists" {
		return ErrKeyExists
	}
	return err
}

func (s *KVServiceImpl) Get(key string) (interface{}, error) {
	value, err := s.repo.Get(key)
	if err != nil && err.Error() == "key not found" {
		return nil, ErrKeyNotFound
	}
	return value, err
}

func (s *KVServiceImpl) Update(key string, value interface{}) error {
	_, err := s.repo.Get(key)
	if err != nil {
		return ErrKeyNotFound
	}
	return s.repo.Update(key, value)
}

func (s *KVServiceImpl) Delete(key string) error {
	_, err := s.repo.Get(key)
	if err != nil {
		return ErrKeyNotFound
	}
	return s.repo.Delete(key)
}
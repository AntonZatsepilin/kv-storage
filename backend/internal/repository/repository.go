package repository

import "github.com/tarantool/go-tarantool/v2"

type KeyValueRepository interface {
	SetValue(key, value string) error
	GetValueByKey(key string) (string, error)
	UpdateValue(key, value string) error
	DeleteValue(key string) error
}

type Repository struct {
	KeyValueRepository
}

func NewRepository(db *tarantool.Connection) *Repository {
	return &Repository{
		KeyValueRepository: NewKeyValueRepository(db),
	}
}
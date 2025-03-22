package repository

import "github.com/tarantool/go-tarantool"

type KVRepository interface {
	Create(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Update(key string, value interface{}) error
	Delete(key string) error
}

type Repository struct {
	KVRepository
}

func NewRepository(conn *tarantool.Connection) *Repository {
	return &Repository{
		KVRepository: NewKVTarantool(conn),
	}
}
package repository

import (
	"github.com/tarantool/go-tarantool/v2"
)

type KeyValueTarantool struct {
	db *tarantool.Connection
}

func NewKeyValueRepository(db *tarantool.Connection) *KeyValueTarantool {
	return &KeyValueTarantool{db: db}
}
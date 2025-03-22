package repository

import (
	"fmt"

	"github.com/tarantool/go-tarantool"
)

type KVTarantool struct {
	conn *tarantool.Connection
}

func NewKVTarantool(conn *tarantool.Connection) *KVTarantool {
	return &KVTarantool{conn: conn}
}

func (r *KVTarantool) Create(key string, value interface{}) error {
	_, err := r.conn.Insert("kv", []interface{}{key, value})
	if err != nil && err.Error() == "Duplicate key exists in unique index 'primary' in space 'kv'" {
		return fmt.Errorf("key already exists")
	}
	return err
}

func (r *KVTarantool) Get(key string) (interface{}, error) {
	resp, err := r.conn.Select("kv", "primary", 0, 1, tarantool.IterEq, []interface{}{key})
	if err != nil {
		return nil, err
	}
	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("key not found")
	}
	return resp.Data[0].([]interface{})[1], nil
}

func (r *KVTarantool) Update(key string, value interface{}) error {
	_, err := r.conn.Update("kv", "primary", []interface{}{key}, []tarantool.Op{
		{"=", 1, value},
	})
	return err
}

func (r *KVTarantool) Delete(key string) error {
	_, err := r.conn.Delete("kv", "primary", []interface{}{key})
	return err
}
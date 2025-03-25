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

func (r *KeyValueTarantool) SetValue(key, value string) error {
    _, err := r.db.Do(tarantool.NewInsertRequest("kv").Tuple([]interface{}{key, value}),).Get()
    return err
}

func (r *KeyValueTarantool) GetValueByKey(key string) (string, error) {
    resp, err := r.db.Do(tarantool.NewSelectRequest("kv").Limit(1).Iterator(tarantool.IterEq).Key([]interface{}{key}),).Get()
    if err != nil {
        return "", err
    }
	
    if len(resp) == 0 {
        return "", nil
    }

    tuple, ok := resp[0].([]interface{})
    if !ok || len(tuple) < 2 {
        return "", nil
    }
	
    value, ok := tuple[1].(string)
    if !ok {
        return "", nil
    }
    return value, nil
}

func (r *KeyValueTarantool) UpdateValue(key, value string) error {
	_, err := r.db.Do(tarantool.NewUpdateRequest("kv").Key([]interface{}{key}).Operations(tarantool.NewOperations().Assign(1, value),),).Get()

    return err
}
package repository

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/tarantool/go-tarantool/v2"
)

var (
    ErrKeyExists    = errors.New("key already exists")
    ErrKeyNotFound  = errors.New("key not found")
    ErrInvalidData  = errors.New("invalid data format")
)

type KeyValueTarantool struct {
	db *tarantool.Connection
}

func NewKeyValueRepository(db *tarantool.Connection) *KeyValueTarantool {
	return &KeyValueTarantool{db: db}
}

func (r *KeyValueTarantool) SetValue(key, value string) error {

        log := logrus.WithFields(logrus.Fields{
        "operation": "set_value",
        "key":       key,
    })

    _, err := r.db.Do(tarantool.NewInsertRequest("kv").Tuple([]interface{}{key, value}),).Get()

    if err != nil {
        log.WithError(err).Error("failed to set value")
        return ErrKeyExists
    }

    log.Debug("value set successfully")
    return nil
}

func (r *KeyValueTarantool) GetValueByKey(key string) (string, error) {
    log := logrus.WithFields(logrus.Fields{
        "operation": "get_value",
        "key":       key,
    })
    resp, err := r.db.Do(tarantool.NewSelectRequest("kv").Limit(1).Iterator(tarantool.IterEq).Key([]interface{}{key}),).Get()
    if err != nil {
        log.WithError(err).Error("failed to get value")
        return "", err
    }
	
    if len(resp) == 0 {
        log.Debug("key not found")
        return "", ErrKeyNotFound
    }

    tuple, ok := resp[0].([]interface{})
    if !ok || len(tuple) < 2 {
        log.Error("invalid data format")
        return "", ErrInvalidData
    }
	
    value, ok := tuple[1].(string)
    if !ok {
        return "", ErrInvalidData
    }

    log.Debug("value retrieved successfully")
    return value, nil
}

func (r *KeyValueTarantool) UpdateValue(key, value string) error {
    log := logrus.WithFields(logrus.Fields{
        "operation": "update_value",
        "key":       key,
    })
	resp, err := r.db.Do(tarantool.NewUpdateRequest("kv").Key([]interface{}{key}).Operations(tarantool.NewOperations().Assign(1, value),),).Get()
    if err != nil {
        log.WithError(err).Error("failed to update value")
        return err
    }

    if len(resp) == 0 {
        log.Debug("key not found")
        return ErrKeyNotFound
    }
    
    log.Debug("value updated successfully")
    return nil
}

func (r *KeyValueTarantool) DeleteValue(key string) error {
    log := logrus.WithFields(logrus.Fields{
        "operation": "delete_value",
        "key":       key,
    })
    resp, err := r.db.Do(tarantool.NewDeleteRequest("kv").Key([]interface{}{key}),).Get()

    if err != nil {
        log.WithError(err).Error("failed to delete value")
        return err
    }

    if len(resp) == 0 {
        log.Debug("key not found")
        return ErrKeyNotFound
    }

    log.Debug("value deleted successfully")
    return nil
}

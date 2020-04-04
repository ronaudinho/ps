package ps

import (
	"errors"
	"fmt"
)

var (
	ErrKeyExists       = func(s ...string) error { return errors.New(fmt.Sprintf("key %s exists", s)) }
	ErrKeyDoesNotExist = func(s ...string) error { return errors.New(fmt.Sprintf("key %s does not exist", s)) }
)

// Store is an interface for CRUD store
// in this case, we only C(reate)R(ead) hence U(pdateD(elete) is not required by this interface
type Store interface {
	Add(string, string, interface{}) error
	Get(...string) (interface{}, error)
}

// KVStore is a simple implementation of Store interface
type KVStore map[string]map[string]interface{}

// Add adds key value to a table
// it creates a new table if table does not exist
// it errors if key already exists on table
func (kv KVStore) Add(t, k string, v interface{}) error {
	_, ok := kv[t]
	if !ok {
		kv[t] = make(map[string]interface{})
	}
	_, ok = kv[t][k]
	if ok {
		return ErrKeyExists([]string{t, k}...)
	}
	kv[t][k] = v
	return nil
}

// Get gets values from store
// it traverses store keys included in the input []string
func (kv KVStore) Get(s ...string) (interface{}, error) {
	m := make(map[string]interface{})
	for i, ss := range s {
		if i == 0 {
			_, ok := kv[ss]
			if !ok {
				return nil, ErrKeyDoesNotExist(s...)
			}
			m = kv[ss]
			continue
		}
		mm, ok := m[ss].(map[string]interface{})
		if ok {
			m = mm
			continue
		}
		i, ok := m[ss].(interface{})
		if ok {
			return i, nil
		}
		return nil, ErrKeyDoesNotExist(s...)
	}
	return m, nil
}

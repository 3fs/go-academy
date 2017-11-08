package kvMemory

import (
	"errors"
)

// ErrNotFound is returned when an nonexisting key is requested
var ErrNotFound = errors.New("Key not found")

// KVStorage describes basic functions required for KV storage
type KVStorage interface {
	Add(string, string) error
	Get(string) (string, error)
	Remove(string) error
}

// KV hold all the data in memory
type KV struct {
	data map[string]string
}

// New returns an initialized instance of KV
func New() *KV {
	return &KV{map[string]string{}}
}

// Add writes a given value under the given key
func (kv *KV) Add(k, v string) error {
	kv.data[k] = v
	return nil
}

// Get returns value linked to given key. Returns ErrNotFound when key does not
// exist.
func (kv *KV) Get(k string) (string, error) {
	if v, ok := kv.data[k]; ok {
		return v, nil
	}

	return "", ErrNotFound
}

// Remove deletes the given key from the map. Returns ErrNotFound if key does
// not exist.
func (kv *KV) Remove(k string) error {
	if _, ok := kv.data[k]; !ok {
		return ErrNotFound
	}

	delete(kv.data, k)
	return nil
}

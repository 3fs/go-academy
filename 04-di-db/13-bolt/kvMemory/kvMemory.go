package kvMemory

import (
	"errors"

	bolt "github.com/coreos/bbolt"
)

// ErrNotFound is returned when an nonexisting key is requested
var ErrNotFound = errors.New("Key not found")
var kvBucket = []byte("kv")

// KVStorage describes basic functions required for KV storage
type KVStorage interface {
	Add(string, string) error
	Get(string) (string, error)
	Remove(string) error
}

// KV holds all the data in memory
type KV struct {
	db *bolt.DB
}

// New returns an initialized instance of KV
func New(path string) (*KV, error) {
	b, err := bolt.Open(path, 0644, nil)
	return &KV{db: b}, err
}

// Add writes a given value under the given key
func (kv *KV) Add(k, v string) error {
	return kv.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(kvBucket)
		if err != nil {
			return err
		}

		return b.Put([]byte(k), []byte(v))
	})
}

// Get returns value linked to given key. Returns ErrNotFound when key does not
// exist.
func (kv *KV) Get(k string) (string, error) {
	var val string
	err := kv.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(kvBucket)
		val = string(b.Get([]byte(k)))
		return nil
	})
	return val, err
}

// Remove deletes the given key from the map. Returns ErrNotFound if key does
// not exist.
func (kv *KV) Remove(k string) error {
	return kv.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(kvBucket)
		return b.Delete([]byte(k))
	})
}

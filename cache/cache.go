package cache

import (
	lru "github.com/hashicorp/golang-lru"
	"sync"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
	Close() error
}

type StorageWithCache struct {
	store Storage
	cache *lru.Cache
	lock  sync.RWMutex
}

func NewCache(store Storage, cacheSize int) (Storage, error) {
	r := &StorageWithCache{}
	r.cache, _ = lru.New(cacheSize)
	r.store = store
	return r, nil
}

func (s *StorageWithCache) Get(key string) (string, error) {
	// local cache get
	c, exist := s.cache.Get(key)
	if exist {
		return c.(string), nil
	}
	s.lock.Lock()
	v, err := s.store.Get(key)
	s.lock.Unlock()
	if v != "" {
		// grab data from redis and write into local cache
		s.cache.Add(key, v)
	}
	return v, err
}

func (s *StorageWithCache) Set(key string, value string) error {
	old, exist := s.cache.Get(key)
	s.cache.Add(key, value)
	s.lock.Lock()
	err := s.store.Set(key, value)
	s.lock.Unlock()
	if err != nil && exist {
		s.cache.Add(key, old)
	}
	return err
}

func (s *StorageWithCache) Delete(key string) error {
	old, exist := s.cache.Get(key)

	s.cache.Remove(key)
	s.lock.Lock()
	err := s.store.Delete(key)
	s.lock.Unlock()

	if err != nil && exist {
		s.cache.Add(key, old)
	}
	return err
}

func (s *StorageWithCache) Close() error {
	return s.store.Close()
}


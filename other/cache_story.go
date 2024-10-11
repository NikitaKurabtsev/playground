package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type KVDatabase interface {
	Get(string) (string, error)
	GetKeys() ([]string, error)
}

type RedisDatabase struct {
	// ...
}

type Cache struct {
	cache      map[string]CacheItem
	db         KVDatabase
	mu         sync.RWMutex
	expiration time.Duration
}

type CacheItem struct {
	value      string
	expiration time.Time
}

func NewCache(db KVDatabase, expiration time.Duration) *Cache {
	return &Cache{
		cache:      make(map[string]CacheItem),
		db:         db,
		expiration: expiration,
	}
}

func (c *Cache) Get(key string) (string, error) {
	c.mu.RLock()
	item, found := c.cache[key]
	c.mu.RUnlock()

	if found && time.Now().Before(item.expiration) {
		return item.value, nil
	}

	value, err := c.db.Get(key)
	if err != nil {
		return "", err
	}

	c.mu.Lock()
	c.cache[key] = CacheItem{
		value:      value,
		expiration: time.Now().Add(c.expiration),
	}
	c.mu.Unlock()

	return value, nil
}

func (d *RedisDatabase) Get(key string) (string, error) {
	/// ...
	return key, errors.New("not implemented")
}

func (d *RedisDatabase) GetKeys() ([]string, error) {
	// ...
	return nil, errors.New("not implemented")
}

func main() {
	redisDB := &RedisDatabase{}
	cache := NewCache(redisDB, 5*time.Minute)

	value, err := cache.Get("key_value")
	if err != nil {
		fmt.Printf("Error getting the value %v\n", err)
		time.Sleep(2 * time.Second)
	}

	value, err = cache.Get("key_value")
	if err != nil {
		fmt.Printf("Error getting the value %v\n", err)
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("value %v", value)
}

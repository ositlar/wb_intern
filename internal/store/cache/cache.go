package cache

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/ositlar/wb_intern_1/internal/store/sqlstore"
)

type Cache struct {
	data  map[string]string
	mutex sync.Mutex
}

func NewCache(store *sqlstore.Store) *Cache {
	cache := &Cache{
		data: make(map[string]string),
	}
	orders, _ := store.SelectAll()
	var (
		id   string
		info map[string]interface{}
	)
	for i := 0; i < len(orders); i++ {
		id = orders[i].Id
		info = orders[i].Info
		cache.Set(id, info)
	}
	return cache
}

func (c *Cache) Set(key string, value map[string]interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	data, err := json.Marshal(value)
	if err != nil {
		log.Fatal("Method cache.Set: ", err)
	}
	c.data[key] = string(data)
}

func (c *Cache) Get(key string) (map[string]interface{}, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	jsonData := make(map[string]interface{})
	err := json.Unmarshal([]byte(c.data[key]), &jsonData)
	if err != nil {
		return nil, errors.New("Cache get error")
	}
	return jsonData, nil
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}

func (c *Cache) GetAll() map[string]string {
	return c.data
}

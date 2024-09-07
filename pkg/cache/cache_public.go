package cache

import (
	"errors"
	"time"
)

func New() *Cache {
	cache := Cache{
		items: make(map[string]Item),
	}

	return &cache
}

func (c *Cache) Get(key string) (interface{}, bool) {

	c.RLock()

	item, found := c.items[key]

	// ключ не найден
	if !found {
		return nil, false
	}

	// Если в момент запроса кеш устарел возвращаем nil
	if time.Now().UnixNano() > int64(item.LifeCycle) {
		defer c.Delete(key)
		defer c.RUnlock()

		return nil, true
	}

	defer c.RUnlock()
	return item.Value, false
}

func (c *Cache) Set(key string, value interface{}, lifeTime time.Duration) {
	c.Lock()

	defer c.Unlock()

	item := Item{
		Value:     value,
		LifeCycle: time.Now().UnixNano() + int64(lifeTime),
	}

	c.items[key] = item
}

func (c *Cache) Delete(key string) error {

	c.Lock()

	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("Key not found")
	}

	delete(c.items, key)

	return nil
}

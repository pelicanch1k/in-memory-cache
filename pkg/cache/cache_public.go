package cache

import (
	"time"
	"errors"
)

func New() *Cache {
    cache := Cache{
        items: make(map[string]Item),
    }

    return &cache
}

func (c *Cache) Get(key string) (interface{}, bool) {

    c.RLock()

    defer c.RUnlock()

    item, found := c.items[key]

    // ключ не найден
    if !found {
        return nil, false
    }

    // Проверка на установку времени истечения, в противном случае он бессрочный
    if item.Expiration > 0 {

        // Если в момент запроса кеш устарел возвращаем nil
        if time.Now().UnixNano() > item.Expiration {
            return nil, true
        }

    }

    return item.Value, false
}

func (c *Cache) Set(key string, value interface{}) {

    var expiration int64 = 0

    c.Lock()

    defer c.Unlock()

    c.items[key] = Item{
        Value:      value,
        Expiration: expiration,
        Created:    time.Now(),
    }

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
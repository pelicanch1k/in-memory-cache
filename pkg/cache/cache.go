package cache

import (
	"sync"
	"time"
)

type Item struct {
    Value      interface{}
    Created    time.Time
    Expiration int64
}

type Cache struct {
    sync.RWMutex
    items  map[string]Item
}
package cache

import (
	"sync"
)

type Item struct {
	Value     interface{}
	LifeCycle int64
}

type Cache struct {
	sync.RWMutex
	items map[string]Item
}

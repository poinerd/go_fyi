package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}
func NewCache() *Cache { //To avoid goo= copying stuff
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) SetCache(key, value string) { 
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) GetCache(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, exists := c.data[key]
	return val, exists
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			key := fmt.Sprintf("user_%d", id)
			value := fmt.Sprintf("data_%d", id)

			cache.SetCache(key, value)
		}(i)
	}

	wg.Wait()

	if val, found := cache.GetCache("user_2"); found {
		fmt.Println("found value:", val)
	} else {
		fmt.Println("key not found")
	}
}
package main

import (
	"fmt"
	"sync"
)

// Cache represents our thread-safe, in-memory key-value store.
type Cache struct {
	// mu protects the underlying map from concurrent read/write conflicts.
	mu   sync.RWMutex
	data map[string]string
}

// NewCache initializes and returns a new Cache.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Set adds or updates a key-value pair. 
// It uses a Lock() because writing requires exclusive access.
func (c *Cache) Set(key, value string) {
	c.mu.Lock()         // Acquire exclusive write lock
	defer c.mu.Unlock() // Ensure lock is released when the function returns
	
	c.data[key] = value
}

// Get retrieves a value by its key.
// It uses an RLock() because multiple goroutines can safely read at the same time
// as long as no one is writing.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()         // Acquire shared read lock
	defer c.mu.RUnlock() // Ensure lock is released
	
	val, exists := c.data[key]
	return val, exists
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup

	// Simulating concurrent writes
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("user_%d", id)
			value := fmt.Sprintf("data_%d", id)
			cache.Set(key, value)
		}(i)
	}

	wg.Wait() // Wait for all writes to finish

	// Safely reading a value
	if val, found := cache.Get("user_2"); found {
		fmt.Println("Found value:", val)
	} else {
		fmt.Println("Key not found")
	}
}
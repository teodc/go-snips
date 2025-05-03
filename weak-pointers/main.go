package main

import (
	"fmt"
	"runtime"
	"weak"
)

// Weak pointers are basically a way to reference a chunk of memory without locking it down
// so that the garbage collector can free it when no strong pointers are referencing it anymore.
// When the referenced chunk of memory is freed, the weak pointer will be set to nil.
// Best practices:
// - Use weak pointers to avoid memory leaks
// - Always check if a weak pointer is nil before dereferencing it
// - Don't overuse them

type Cache struct {
	items map[string]weak.Pointer[string]
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]weak.Pointer[string]),
	}
}

func (c *Cache) Add(key string, item *string) {
	// Reference the item with a weak pointer
	wp := weak.Make(item)
	c.items[key] = wp
}

func (c *Cache) Get(key string) (*string, bool) {
	// Get the weak pointer from the cache
	wp, exists := c.items[key]
	if !exists {
		return nil, false
	}

	// Check if the weak pointer is still valid
	if p := wp.Value(); p != nil {
		return p, true
	}

	// The weak pointer is no longer valid, remove it from the cache
	delete(c.items, key)

	return nil, false
}

func main() {
	c := NewCache()

	// Utility function to print the current memory usage
	printMemoryUsage := func() {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Memory usage: %.2f MB\n", float64(m.Alloc)/1024/1024)
	}

	// Utility function to create a 10 MB string
	createBigString := func() *string {
		b := make([]byte, 10<<20)
		s := string(b)
		return &s
	}

	// Add a big string to the cache
	bs := createBigString()
	c.Add("bs", bs)
	if _, exists := c.Get("bs"); exists {
		fmt.Println("Big string found in cache")
	}

	// Free the big string
	bs = nil

	fmt.Println("Before GC:")
	printMemoryUsage()
	runtime.GC() // Run the garbage collector
	fmt.Println("After GC:")
	printMemoryUsage()

	// Check if the big string is still in the cache
	if item, exists := c.Get("bs"); exists {
		fmt.Printf("Still holding %.2f MB in cache\n", float64(len(*item))/1024/1024)
	} else {
		fmt.Println("Big string not found in cache")
	}
}

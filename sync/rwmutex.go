package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	data    map[string]string
	rwmutex sync.RWMutex
}

func (sm *SafeMap) Get(key string) (string, bool) {
	sm.rwmutex.RLock()
	defer sm.rwmutex.RUnlock()

	value, exists := sm.data[key]

	fmt.Printf("Reading %s: %s\n", key, value)
	time.Sleep(100 * time.Millisecond)
	return value, exists
}

func (sm *SafeMap) Set(key, value string) {
	sm.rwmutex.Lock()
	defer sm.rwmutex.Unlock()

	fmt.Printf("Writing %s: %s\n", key, value)
	sm.data[key] = value
	time.Sleep(100 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	var sm = SafeMap{
		data: make(map[string]string),
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		sm.Set("apple", "red")
	}()

	go func() {
		defer wg.Done()
		sm.Set("banana", "yellow")
	}()

	go func() {
		defer wg.Done()
		sm.Set("watermelon", "green")
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			sm.Get("banana")
			sm.Get("apple")
		}(i)
	}

	wg.Wait()
	fmt.Println("All done mate!")
}

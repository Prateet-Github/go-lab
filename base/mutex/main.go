package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock() // lock

	c.count++

	c.mutex.Unlock() // unlock
}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			counter.Increment()
		}()
	}

	wg.Wait()

	fmt.Println("Final count:", counter.count)
}

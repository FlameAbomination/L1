package main

import (
	"fmt"
	"os"
	"sync"
)

type Counter struct {
	lock    sync.Mutex
	counter int
}

func UnsafeCounter(n int) {
	var wg sync.WaitGroup
	var counter Counter
	counter.counter = 0

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			counter.counter++
			defer wg.Done()
		}()
	}

	wg.Wait()
	fmt.Fprintln(os.Stdout, counter.counter)
}

func SafeCounter(n int) {
	var wg sync.WaitGroup
	var counter Counter
	counter.counter = 0

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			counter.lock.Lock()
			counter.counter++
			counter.lock.Unlock()
			defer wg.Done()
		}()
	}

	wg.Wait()
	fmt.Fprintln(os.Stdout, counter.counter)
}

func task18() {
	UnsafeCounter(250)
	SafeCounter(250)
}

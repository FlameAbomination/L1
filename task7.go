package main

import (
	"fmt"
	"os"
	"sync"
)

type Database struct {
	lock sync.Mutex
	data map[int]int
}

func task7() {
	var wg sync.WaitGroup
	var databaseSync sync.Map
	database := Database{
		data: make(map[int]int),
	}
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				database.lock.Lock()
				database.data[j]++
				database.lock.Unlock()
			}
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				databaseSync.Store(i*5+j, j)
			}
		}(i)
	}

	wg.Wait()
	fmt.Fprintln(os.Stdout, database.data)
	databaseSync.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true
	})
}

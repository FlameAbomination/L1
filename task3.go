package main

import (
	"fmt"
	"os"
	"sync"
)

func task3() {
	var wg sync.WaitGroup

	// Отсутствие Mutex может не оказать влияния на конечный результат,
	// однако это вызвано небольшим количеством элементов
	// При увеличении количества элементов в 100 раз ошибка становится постоянной
	var lock sync.Mutex

	arr := []int{2, 4, 6, 8, 10}

	total := 0
	for _, val := range arr {
		wg.Add(1)
		go func(v int) {
			lock.Lock()
			total += v * v
			lock.Unlock()
			defer wg.Done()
		}(val)
	}

	wg.Wait()
	fmt.Fprintln(os.Stdout, total)
}

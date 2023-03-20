package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func Task5() {
	var wg sync.WaitGroup
	fmt.Fprintln(os.Stdout, "Введите время работы программы:")
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		n = 5
		fmt.Fprintln(os.Stdout, "Неверный формат ввода, будет использовано стандартное значение")
	}
	channel := make(chan int)
	done := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(channel)
				fmt.Fprintln(os.Stdout, "Close producer")
				return
			default:
				channel <- 2
				time.Sleep(20 * time.Millisecond)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			_, ok := <-channel
			if !ok {
				fmt.Fprintln(os.Stdout, "Close consumer")
				return
			}
			time.Sleep(20 * time.Millisecond)
		}
	}()
	time.Sleep(time.Duration(n) * time.Second)
	done <- true

	wg.Wait()
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Task4() {
	var wg sync.WaitGroup
	fmt.Fprintln(os.Stdout, "Введите число воркеров:")

	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		n = 5
		fmt.Fprintln(os.Stdout, "Неверный формат ввода, будет использовано стандартное значение")
	}
	channel := make(chan int)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	wg.Add(1)
	go func() {
		i := 2
		defer wg.Done()
		for {
			select {
			case <-signalChan:
				close(channel)
				fmt.Fprintln(os.Stdout, "Close producer")
				return
			default:
				channel <- i
				i++
				time.Sleep(20 * time.Millisecond)
			}
		}
	}()
	// Для завершения воркеров был выбран способ с закрытием канала, так как
	// все воркеры работают с этим с каналом
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				val, ok := <-channel
				if !ok {
					fmt.Fprintln(os.Stdout, "Close worker")
					return
				}
				fmt.Fprintln(os.Stdout, val)
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}

	wg.Wait()
}

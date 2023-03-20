package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func task4() {
	var wg sync.WaitGroup
	n := 5
	channel := make(chan int)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-signalChan:
				close(channel)
				fmt.Fprintln(os.Stdout, "Close producer")
				return
			default:
				channel <- 2
				time.Sleep(20 * time.Millisecond)
			}
		}
	}()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				_, ok := <-channel
				if !ok {
					fmt.Fprintln(os.Stdout, "Close consumer")
					return
				}
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}

	wg.Wait()
}

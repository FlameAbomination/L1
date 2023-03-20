package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func Task6() {
	var wg sync.WaitGroup

	quit := make(chan bool)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Fprintln(os.Stdout, "Stopped")
				wg.Done()
				return
			default:
				time.Sleep(25 * time.Millisecond)
			}
		}
	}()

	ch := make(chan string)
	wg.Add(1)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("Stopped 2")
				wg.Done()
				return
			}
			fmt.Println(v)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopped 3")
				wg.Done()
				return
			default:
				time.Sleep(25 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(250 * time.Millisecond)
	quit <- true
	close(ch)
	cancel()
	wg.Wait()
}

package main

import (
	"fmt"
	"os"
	"time"
)

func sleep(seconds time.Duration) {
	if <-time.After(seconds); true {
		return
	}
}

func sleep2(seconds time.Duration) {
	start := time.Now()
	end := time.Now()
	for end.Sub(start) < seconds {
		end = time.Now()
	}
}

func task25() {
	fmt.Fprintln(os.Stdout, time.Now())
	sleep(time.Duration(1) * time.Second)
	fmt.Fprintln(os.Stdout, time.Now())
	sleep2(time.Duration(1) * time.Second)
	fmt.Fprintln(os.Stdout, time.Now())
}

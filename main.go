package main

import (
	"sync/atomic"
	"time"
)

func main() {
	var progress uint32
	done := make(chan struct{})
	defer close(done)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-time.After(24 * time.Hour):
				println("Hết tết rồi", atomic.LoadUint32(&progress))
				atomic.AddUint32(&progress, 1)
			}
		}
	}()

	time.Sleep(3 * 24 * time.Hour)
}

package main

import (
	"sync/atomic"
	"time"
)

func main() {
	var progress uint32
	done := make(chan interface{})
	defer close(done)
	go func() {
		for {
			select {
			case <-done:
				println("Hết tết thật rồi!")
				return
			case <-time.After(time.Second):
				atomic.AddUint32(&progress, 1)
				println("Sắp hết tết rồi", atomic.LoadUint32(&progress))
			}
		}
	}()

	time.Sleep(4 * time.Second)
	done <- true
	time.Sleep(2 * time.Second)
}

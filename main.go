package main

import (
	"sync/atomic"
	"time"
)

func main() {
	var progress uint32
	done := make(chan interface{}, 1)
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
				// Mùng 3 là hết tết :)
				if atomic.LoadUint32(&progress) == 3 {
					done <- true
				}
			}
		}
	}()

	time.Sleep(4 * time.Second)
}

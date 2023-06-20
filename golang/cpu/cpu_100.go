package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			for {
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

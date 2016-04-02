package main

import "sync"

func main() {
	for i := 0; i < 10; i++ {
		var wg sync.WaitGroup

		intChan := make(chan int)

		for j := 0; j < 10; j++ {
			wg.Add(1)
			go ole(intChan)
		}

		go func() {
			for {
				select {
				case ole := <-intChan:
					println(ole)
					wg.Done()
				}
			}
		}()

		wg.Wait()
	}
}

func ole(pool chan int) {
	pool <- 1
}

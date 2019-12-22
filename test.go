package main

import (
	"fmt"
	"time"
)

func main() {
	signalChn := make(chan struct{})

	for i := 0; i < 5; i++ {
		go func(n int) {
		loop:
			for {
				time.Sleep(1 * time.Second)
				select {
				case <-signalChn:
					fmt.Printf("%dth received signal~\n", n)
					break loop
				default:
					fmt.Printf("%dth goroutine are waiting signal~\n", n)
				}
			}
		}(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("send signal~")
	// signalChn <- struct{}{}
	close(signalChn)
	time.Sleep(10 * time.Second)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go lang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	// READ Only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh

		fmt.Println("is channel open", isChannelOpen)
		fmt.Println("value from channel", val)


		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	// SEND only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// myCh <- 42
		// myCh <- 41
		myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}

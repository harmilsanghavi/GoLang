package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	//Recive Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh

		fmt.Println(val)
		fmt.Println(isChanelOpen)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	//Send Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}

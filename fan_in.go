package main

import (
	"fmt"
	"sync"
)

func fanIn() {

	var inputChan [6]chan int

	for i := 0; i < 6; i++ {
		inputChan[i] = make(chan int)
	}
	outputChan := make(chan int, 7)

	wg := sync.WaitGroup{}
	wg.Add(6)
	for i := 0; i < 6; i++ {
		go processInput(inputChan[i], outputChan, &wg)
	}

	go func() {

		wg.Wait()
		close(outputChan)
	}()

	for i := 0; i < 6; i++ {
		inputChan[i] <- i + 1
		close(inputChan[i])
	}

	for ans := range outputChan {
		fmt.Printf("final value is %d \n", ans)
	}
}

func processInput(inputChannel chan int, outputChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	x := <-inputChannel
	outputChan <- x
}

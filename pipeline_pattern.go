package main

import (
	"fmt"
	"sync"
)

func processInput(input chan int, output chan int, worker int, wg *sync.WaitGroup) {

	defer wg.Done()
	for in := range input {
		output <- in * worker
		fmt.Printf("%d input process by %d go routine \n", in, worker)
	}
}

func pipelinePattern(n int) {

	// input channel
	chaninput := make(chan int, 50)

	chanoutput := make(chan int, 50)

	// call n goroutines to put the input into input channel

	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}
	for i := 1; i <= n; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			chaninput <- i
			fmt.Printf("%d value added by go routine %d \n", i, i)
		}()
	}

	go func() {
		wg1.Wait()
		close(chaninput)
	}()

	for i := 1; i <= n; i++ {
		wg2.Add(1)
		go processInput(chaninput, chanoutput, i, &wg2)
	}

	go func() {
		wg2.Wait()
		close(chanoutput)
	}()

	for out := range chanoutput {
		fmt.Printf("values are %d\n", out)
	}

	// call n goroutines to take the value from input channel and process it and put the result to output channel

	// retrive all the values from out put channel and print them

}

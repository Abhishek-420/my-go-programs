package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func PingpongpatternWithContext() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

	defer cancel()

	wg := sync.WaitGroup{}
	first := make(chan struct{})
	second := make(chan struct{})

	wg.Add(2)
	go printFirstNameWithContext(first, second, ctx, &wg)
	go printLastNameWithContext(first, second, ctx, &wg)

	first <- struct{}{}

	wg.Wait()

	fmt.Println("go routine successfully executed")

}

func printFirstNameWithContext(first chan struct{}, second chan struct{}, ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()
	for {

		select {
		case <-ctx.Done():
			return
		case <-first:
		}

		fmt.Println("Abhishek")

		select {
		case <-ctx.Done():
			return
		case second <- struct{}{}:
		}

	}
}

func printLastNameWithContext(first chan struct{}, second chan struct{}, ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()
	for {

		select {
		case <-ctx.Done():
			return
		case <-second:
		}

		fmt.Println("Rajput")

		select {
		case <-ctx.Done():
			return
		case first <- struct{}{}:
		}
	}
}

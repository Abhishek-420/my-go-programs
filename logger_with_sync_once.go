package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	l *log.Logger
}

var (
	instance *Logger
	once     sync.Once
)

func getLogger(i int) *Logger {

	once.Do(func() {
		fmt.Printf("logger successfully created by %d \n", i)
		instance = &Logger{
			l: log.New(os.Stdout, "LOG :", log.Ldate),
		}

	})
	return instance
}

func logger(n int) {

	wg := sync.WaitGroup{}

	for i := 1; i <= n; i++ {

		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			log := getLogger(i)
			log.l.Printf("go routine  %d run successfully ", id)
		}(i)

	}

	wg.Wait()
}

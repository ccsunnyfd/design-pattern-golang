package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2)))
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

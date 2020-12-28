package main

import (
	"fmt"
	"time"
)

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 5)
	}
}

func main() {
	// create channel to comunication
	channel := make(chan int)

	// create workers (threads)
	for i := 1; i <= 5; i++ {
		go worker(channel)
	}

	// pass values to worker by channel
	for i := 0; i < 100; i++ {
		channel <- i
	}
}

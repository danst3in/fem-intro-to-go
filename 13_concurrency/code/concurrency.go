package main

import (
	"fmt"
	"sync"
	"time"
)

// func say(s string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(s)
// 		time.Sleep(time.Millisecond * 300)
// 	}
// }

var wg sync.WaitGroup

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("PANIC")
	}
}

func printStuff() {
	defer wg.Done()
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	wg.Add(1)
	go printStuff()
	wg.Wait()

	// go say("Hello")
	// say("There")
}

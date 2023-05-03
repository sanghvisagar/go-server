package main

import (
	"fmt"
	"time"
)

func main() {
	// msgChannel := make(chan string)

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	msgChannel <- "Hello"
	// }()

	// fmt.Println("Msg from goroutine", <-msgChannel)

	// go func() {
	// 	fmt.Println("2nd Goroutine")
	// }()

	// msgChannel := make(chan []string, 2)

	// go func() {
	// 	msgChannel <- []string{"Hello", "Sagar"}
	// }()

	// go func() {
	// 	msgChannel <- []string{"Sanghvi"}
	// }()

	// fmt.Println(<-msgChannel)
	// fmt.Println(<-msgChannel)

	// done := make(chan bool, 1)
	// go worker(done)
	// <-done

	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)
	// go ping(pings, "playing ping pong")
	// go pong(pongs, pings)

	// fmt.Println(<-pongs)

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		// time.Sleep(2 * time.Second)
		c1 <- "sample"
	}()

	go func() {
		// time.Sleep(2 * time.Second)
		c2 <- "go select"
	}()

	// for i := 0; i < 3; i++ {
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
		// case <-time.After(1 * time.Second):
		// 	fmt.Println("timeout...")
	}
	// }

}

func worker(done chan bool) {
	fmt.Println("working....")
	time.Sleep(time.Second)
	fmt.Println("Done")
	done <- false
}

func ping(ping chan<- string, msg string) {
	ping <- msg
}

func pong(pong chan<- string, ping <-chan string) {
	msg := <-ping
	pong <- msg
}

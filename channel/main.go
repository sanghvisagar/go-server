package main

import "fmt"

func main() {
	msgChannel := make(chan string)

	go func() {
		msgChannel <- "Hello"
	}()

	fmt.Println("Msg from goroutine", <-msgChannel)
}

package main

import (
	"fmt"
	"time"
)

func main() {

	go func(msg string) {
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}("Hello! World,")

	fmt.Println("Hi")

	time.Sleep(2 * time.Second)

}

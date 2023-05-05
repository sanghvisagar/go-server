package main

import "fmt"

func main() {
	cnp := make(chan func(), 10)
	for i := 0; i < 4; i++ {
		fmt.Println("inside for loop")
		go func() {
			for f := range cnp {
				f()
			}
		}()
	}
	cnp <- func() {
		fmt.Println("HERE1")
	}
	fmt.Println("Hello")
}

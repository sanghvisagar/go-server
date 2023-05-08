package main

import "fmt"

func main() {
	cnp := make(chan func(), 10)

	cnp <- func() {
		fmt.Println("HERE1")
	}
	cnp <- func() {
		fmt.Println("HERE1")
	}
	cnp <- func() {
		fmt.Println("HERE1")
	}
	cnp <- func() {
		fmt.Println("HERE1")
	}
	close(cnp)

	for i := 0; i < 4; i++ {
		fmt.Println("inside for loop")
		go func() {
			for f := range cnp {

				f()
			}
		}()
	}

	fmt.Println("Hello")
}

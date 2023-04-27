package main

import "fmt"

func main() {
	const a = 11
	if a%3 == 0 {
		fmt.Println("inside if")
	} else if a%2 == 0 {
		fmt.Println("inside else")
	} else {
		fmt.Println("Inside final else")
	}
}
